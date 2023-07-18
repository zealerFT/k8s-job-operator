package route

import (
	"k8s-job-operator/http/api"
	"k8s-job-operator/http/middleware"

	"github.com/gin-gonic/gin"
)

func AppAPI(opts ...func(engine *gin.Engine)) func(s *gin.Engine) {
	return func(s *gin.Engine) {
		for _, opt := range opts {
			opt(s)
		}

		// 设置静态文件路由
		s.Static("/static", "./frontend/build/static")
		// 处理manifest.json请求
		s.StaticFile("/manifest.json", "./frontend/build/manifest.json")
		s.StaticFile("/favicon.ico", "./frontend/build/favicon.ico")
		s.StaticFile("/logo192.png", "./frontend/build/logo192.png")
		s.StaticFile("/logo512.png", "./frontend/build/logo512.png")
		s.StaticFile("/", "./frontend/build/index.html")

		// Not Found
		s.NoRoute(api.Handle404)
		// Health Check
		s.GET("/check", api.Health)

		r := authRouteGroup(s, "/api/v1/operator")
		r.GET("job/list", api.JobList)
		r.POST("job/operator", api.JobOperator)
		r.POST("job/delete", api.JobDelete)
		r.GET("job/log", api.JobLogsGet)
		r.GET("job/yaml", api.JobYamlGet)
	}
}

// 需要jwt鉴权的group
func authRouteGroup(s *gin.Engine, relativePath string) *gin.RouterGroup {
	group := s.Group(relativePath)
	group.Use(middleware.Auth())
	return group
}
