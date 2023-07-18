package job

import (
	"context"

	"k8s-job-operator/config"
	"k8s-job-operator/model"
	"k8s-job-operator/service/snowflakesvc"

	lru "github.com/hashicorp/golang-lru"
	"github.com/rs/zerolog"
	"go.uber.org/dig"
	"k8s.io/client-go/kubernetes"
)

type Hub struct {
	dig.In

	Job Job
}

//go:generate mockgen -destination mock/job.go -package jobmock k8s-job-operator/biz/job Job
type Job interface {
	JobOperator(ctx context.Context, args, namespace, name, image string) error
	JobList(ctx context.Context, namespace string) ([]*model.Job, error)
	JobDelete(ctx context.Context, namespace, jobName string) error
	JobLogGet(ctx context.Context, namespace, jobName string) (string, error)
	JobYamlGet(ctx context.Context, namespace, jobName string) (string, error)
}

type Jobbiz struct {
	dig.In

	Cfg *config.AppConfig // 配置

	Snowflake *snowflakesvc.Snowflake
	Log       zerolog.Logger
	Lru       *lru.Cache
	K8s       *kubernetes.Clientset
}

// Object 当invoke的时候会实例化Larkbiz对象，并使用go的interface实现方式实现接口，所以这个变量在初始化的时候已是一个全局变量
var Object Job

func NewJob(entry Jobbiz) Job {
	Object = &entry
	return Object
}
