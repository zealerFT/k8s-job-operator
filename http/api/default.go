package api

import (
	"net/http"

	"k8s-job-operator/http/middleware"
	"k8s-job-operator/model"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func Health(c *gin.Context) {
	c.String(http.StatusOK, ":)")
}

func Handle404(c *gin.Context) {
	c.String(http.StatusNotFound, "404 NotFound")
}

func JobList(c *gin.Context) {
	list, err := middleware.Dependency(c).JobHub.Job.JobList(c, c.Query("namespace"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "failed to get Job list"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"list": list})
}

func JobDelete(c *gin.Context) {
	var body struct {
		JobName   string `json:"job_name"`
		Namespace string `json:"namespace"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		log.Err(err).Msg("failed to bind JobDelete body")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "JobDelete malformed request"})
		return
	}
	err := middleware.Dependency(c).JobHub.Job.JobDelete(c, body.Namespace, body.JobName)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "failed to get Job JobDelete"})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func JobOperator(c *gin.Context) {
	body := model.RequestBody{}
	if err := c.ShouldBindJSON(&body); err != nil {
		log.Err(err).Msg("failed to bind JobOperator body")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "JobOperator malformed request"})
		return
	}
	err := middleware.Dependency(c).JobHub.Job.JobOperator(c, body.Args, body.Namespace, body.Name, body.Image)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "failed to get Job Operator"})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func JobLogsGet(c *gin.Context) {
	resp, err := middleware.Dependency(c).JobHub.Job.JobLogGet(c, c.Query("namespace"), c.Query("job_name"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "failed to get Job JobLogsGet"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": resp})
}

func JobYamlGet(c *gin.Context) {
	resp, err := middleware.Dependency(c).JobHub.Job.JobYamlGet(c, c.Query("namespace"), c.Query("job_name"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "failed to get Job yaml"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": resp})
}
