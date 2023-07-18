package dep

import (
	jobBiz "k8s-job-operator/biz/job"
	"k8s-job-operator/config"
	"k8s-job-operator/service/snowflakesvc"

	lru "github.com/hashicorp/golang-lru"
	"go.uber.org/dig"
)

func DIHttpDependency() (out *HttpDependency) {
	container := DI()
	if err := container.Invoke(func(dep HttpDependency) { out = &dep }); err != nil {
		panic(err)
	}

	return
}

type HttpDependency struct {
	dig.In
	Config    *config.AppConfig
	Snowflake *snowflakesvc.Snowflake
	Lru       *lru.Cache

	JobHub jobBiz.Hub
}
