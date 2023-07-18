package job

import (
	"bytes"
	"context"
	"errors"

	"k8s-job-operator/model"
	"k8s-job-operator/util"

	"gopkg.in/yaml.v3"
	v1 "k8s.io/api/batch/v1"
	apiv1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
)

func (j *Jobbiz) JobList(ctx context.Context, namespace string) ([]*model.Job, error) {
	list, err := j.K8s.BatchV1().Jobs(namespace).List(ctx, metav1.ListOptions{})
	if err != nil {
		return []*model.Job{}, err
	}

	containersNames := make([]*model.Job, 0)
	for _, i := range list.Items {
		for _, n := range i.Spec.Template.Spec.Containers {
			// 执行中的任务没有conditions
			status := "executing"
			message := ""
			if len(i.Status.Conditions) > 0 {
				status = string(i.Status.Conditions[0].Type)
				message = i.Status.Conditions[0].Message
			}
			containersNames = append(containersNames, &model.Job{
				JobName:        i.ObjectMeta.Name, // 这里是job的名字，而不是pod的名字
				StartTime:      i.Status.StartTime,
				CompletionTime: i.Status.CompletionTime,
				Active:         i.Status.Active,
				Succeeded:      i.Status.Succeeded,
				Failed:         i.Status.Failed,
				Logs:           n.Name,
				Status:         status,
				Message:        message,
				Namespace:      namespace,
			})
		}
	}

	return containersNames, nil
}

func (j *Jobbiz) JobOperator(ctx context.Context, args, namespace, name, image string) error {
	if namespace == "" {
		namespace = "algorithm"
	}
	jobName := name + "-" + j.Snowflake.Generate().String()
	job := &v1.Job{
		ObjectMeta: metav1.ObjectMeta{
			Name: jobName,
			Labels: map[string]string{
				"app":       "demo",
				"component": "jobs",
				"version":   "default",
			},
			Namespace: namespace,
		},
		Spec: v1.JobSpec{
			BackoffLimit: int32Ptr(0), // backoffLimit 设置为 0，表示在容器失败时不进行重试。
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app":       "demo",
						"component": "jobs",
						"version":   "default",
					},
					Annotations: map[string]string{
						"sidecar.istio.io/inject": "false",
					},
				},
				Spec: apiv1.PodSpec{
					RestartPolicy: "Never", // 设置为 Never，表示容器在失败时不会被重启。
					Containers: []apiv1.Container{
						{
							Name:       name,
							Image:      image,
							Args:       util.StringToSlice(args),
							WorkingDir: "",
							Ports:      []apiv1.ContainerPort{},
							EnvFrom:    nil,
							Env: []apiv1.EnvVar{
								{
									Name:  "APP",
									Value: "demo",
								},
								{
									Name:  "COMPONENT",
									Value: "jobs",
								},
								{
									Name:  "VERSION",
									Value: "default",
								},
							},
							Resources: apiv1.ResourceRequirements{
								Limits: apiv1.ResourceList{
									apiv1.ResourceCPU:    resource.MustParse("50m"),
									apiv1.ResourceMemory: resource.MustParse("50Mi"),
								},
								Requests: apiv1.ResourceList{
									apiv1.ResourceCPU:    resource.MustParse("40m"),
									apiv1.ResourceMemory: resource.MustParse("20Mi"),
								},
							},
						},
					},
				},
			},
		},
	}

	// client
	clientJob := j.K8s.BatchV1().Jobs(namespace)

	// Create job
	j.Log.Info().Msgf("Creating job...")

	create, err := clientJob.Create(ctx, job, metav1.CreateOptions{})
	if err != nil {
		j.Log.Err(err).Msgf("Create job error")
		return err
	}
	j.Log.Info().Msgf("Created job %q.\n", create.GetObjectMeta().GetName())

	return nil
}

func (j *Jobbiz) JobDelete(ctx context.Context, namespace, jobName string) error {
	if namespace == "" {
		namespace = "algorithm"
	}
	client := j.K8s.BatchV1().Jobs(namespace)
	// 只有是完成状态的任务才能删除
	get, err := client.Get(ctx, jobName, metav1.GetOptions{})
	if err != nil {
		j.Log.Error().Err(err).Msgf("get job error")
	}
	// get.Status.Conditions不存在则标识Job还未完成
	if len(get.Status.Conditions) <= 0 {
		j.Log.Error().Msgf("get job status error")
		return errors.New("get job status error")
	}

	j.Log.Info().Msgf("Deleting job...")
	deletePolicy := metav1.DeletePropagationForeground
	if err := client.Delete(ctx, jobName, metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	}); err != nil {
		j.Log.Error().Err(err).Msgf("Delete job error")
	}
	j.Log.Info().Msgf("Deleted job %q.\n", jobName)
	return nil
}

func (j *Jobbiz) JobLogGet(ctx context.Context, namespace, jobName string) (string, error) {
	if namespace == "" {
		namespace = "algorithm"
	}
	client := j.K8s.CoreV1().Pods(namespace)
	labelSelector := metav1.LabelSelector{
		MatchLabels: map[string]string{
			"job-name": jobName,
		},
	}
	podList, err := client.List(ctx, metav1.ListOptions{
		LabelSelector: labels.Set(labelSelector.MatchLabels).String(),
	})
	if podList == nil || len(podList.Items) <= 0 {
		j.Log.Err(err).Msgf("get job log error")
		return "", err
	}
	// Get the logs for the pod
	req := client.GetLogs(podList.Items[0].Name, &apiv1.PodLogOptions{})
	podLogs, err := req.Stream(ctx)
	if err != nil {
		j.Log.Err(err).Msgf("get job log error")
		return err.Error(), err
	}
	defer podLogs.Close()

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(podLogs)
	if err != nil {
		j.Log.Err(err).Msgf("get job log error")
		return "", err
	}
	return buf.String(), nil

}

func (j *Jobbiz) JobYamlGet(ctx context.Context, namespace, jobName string) (string, error) {
	if namespace == "" {
		namespace = "algorithm"
	}

	jobClient := j.K8s.BatchV1().Jobs(namespace)
	job, err := jobClient.Get(ctx, jobName, metav1.GetOptions{})
	if err != nil {
		j.Log.Err(err).Msgf("get job yaml error")
		return err.Error(), err
	}

	job.ManagedFields = nil // delete managedFields
	jobBytes, _ := yaml.Marshal(job)
	str := string(jobBytes)
	return str, nil
}

func int32Ptr(i int32) *int32 { return &i }
