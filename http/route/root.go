package route

import (
	"fmt"
	"strconv"
	"time"

	"k8s-job-operator/util"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
)

func Root(opts ...func(engine *gin.Engine)) func(s *gin.Engine) {
	return func(s *gin.Engine) {
		for _, opt := range opts {
			opt(s)
		}

		// common middleware
		s.Use(
			CORS(),
			util.TimeOffset("x-timezone-offset"),
			util.TimezoneName(),
			Pagination(),
			requestid.New(),
			gzip.Gzip(gzip.DefaultCompression),
		)
	}
}

var defaultCORSConfig = cors.Config{
	AllowAllOrigins: true,
	AllowMethods:    []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS", "HEAD"},
	AllowHeaders: []string{
		"Origin",
		"Accept",
		"Accept-Language",
		"Content-Language",
		"Content-Type",
		"User-Agent",
		"Authorization",
		"x-timezone-offset",
		"x-user-id",
		"X-User-Id",
		"X-Rate-Limit-Token",
		"x-rate-limit-token",
		"X-Timezone-Name",
		"x-timezone-name",
		"X-Lingochamp-Id",
		"x-lingochamp-id",
		"x-user-language",
	},
	AllowCredentials: true,
	MaxAge:           12 * time.Hour,
}

// CORS enable CORS support
func CORS(configs ...cors.Config) gin.HandlerFunc {
	if len(configs) != 0 {
		return cors.New(configs[0])
	}

	return cors.New(defaultCORSConfig)
}

// Pagination sets page, pageSize and pageOffset to *gin.Context
func Pagination() gin.HandlerFunc {
	return func(c *gin.Context) {
		page := getSetItem(c, "page", 1)
		size := getSetItem(c, "pageSize", 20)
		c.Set("pageOffset", (page-1)*size)
		c.Next()
	}
}

func getSetItem(c *gin.Context, k string, d int) int {
	var n int
	if v := c.Query(k); v != "" {
		if i, err := strconv.Atoi(v); err == nil {
			if i > 0 {
				n = i
			}
		}
	}

	if n == 0 {
		n = d
	}

	c.Set(k, n)
	c.Request.URL.Query().Set(k, fmt.Sprintf("%d", n))
	return n
}
