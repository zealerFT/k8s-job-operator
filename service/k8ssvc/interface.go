package k8ssvc

import (
	"k8s-job-operator/config"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func MustNewK8sService(cfg *config.AppConfig) *kubernetes.Clientset {
	conf, err := clientcmd.BuildConfigFromFlags("", cfg.KubeSetting.KubeconfigFile)
	if err != nil {
		panic(err)
	}
	clientset, err := kubernetes.NewForConfig(conf)
	if err != nil {
		panic(err)
	}
	return clientset
}
