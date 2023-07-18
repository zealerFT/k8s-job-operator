package dep

import (
	"log"

	"k8s-job-operator/biz/job"
	"k8s-job-operator/config"
	"k8s-job-operator/service/k8ssvc"
	"k8s-job-operator/service/logsvc"
	"k8s-job-operator/service/lrusvc"
	"k8s-job-operator/service/snowflakesvc"

	"go.uber.org/dig"
)

var Container *dig.Container

// DI 创建容器，注入全局对象
func DI() *dig.Container {
	return NewContainer(
		WithConfig(),
		WithSnowflake(),
		WitchNewLogger(),
		WithLruCache(),
		WithK8s(),
		WithLarkbiz(),
	)
}

func ContainerGet() *dig.Container {
	return Container
}

type Option func(*dig.Container) error

func NewContainer(opts ...Option) *dig.Container {
	container := dig.New()
	for _, opt := range opts {
		if err := opt(container); err != nil {
			log.Fatalf("dig init Container fail: %v", err)
		}
	}
	return container
}

func WithConfig() Option {
	return func(c *dig.Container) error {
		return c.Provide(config.Options)
	}
}

func WitchNewLogger() Option {
	return func(c *dig.Container) error {
		return c.Provide(logsvc.NewLogger)
	}
}

func WithSnowflake() Option {
	return func(c *dig.Container) error {
		return c.Provide(snowflakesvc.MestNewSnowflake)
	}
}

func WithLruCache() Option {
	return func(c *dig.Container) error {
		return c.Provide(lrusvc.NewLruCache)
	}
}

func WithK8s() Option {
	return func(c *dig.Container) error {
		return c.Provide(k8ssvc.MustNewK8sService)
	}
}

/*业务逻辑*/

func WithLarkbiz() Option {
	return func(c *dig.Container) error {
		return c.Provide(job.NewJob)
	}
}
