package http

import (
	"k8s-job-operator/dep"
	"k8s-job-operator/http/middleware"
	"k8s-job-operator/http/route"

	"github.com/gin-gonic/gin"
)

type Server struct {
	*gin.Engine
}

func New(options ...Option) *Server {
	s := &Server{
		Engine: gin.New(),
	}

	for _, option := range options {
		option(s)
	}

	return s
}

// Option 函数式编程
type Option func(*Server)

func ExportLogOption() Option {
	return func(s *Server) {
		s.Engine.Use(gin.LoggerWithConfig(gin.LoggerConfig{SkipPaths: []string{"/metrics", "/check"}}))
	}
}

func SetRouteOption() Option {
	return func(s *Server) {
		route.HTTPServerRoute()(s.Engine)
	}
}

func WithDependency(dep *dep.HttpDependency) Option {
	return func(s *Server) {
		s.Engine.Use(middleware.WithDependency(dep))
	}
}
