package model

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Job struct {
	JobName        string   `gorm:"column:job_name;default:''" json:"job_name"`
	StartTime      *v1.Time `json:"startTime,omitempty"`
	CompletionTime *v1.Time `json:"completionTime,omitempty"`
	Active         int32    `json:"active,omitempty"`
	Succeeded      int32    `json:"succeeded,omitempty"`
	Failed         int32    `json:"failed,omitempty"`
	Logs           string   `json:"logs"`
	Status         string   `json:"status"`
	Message        string   `json:"message"`
	Namespace      string   `json:"namespace"`
}

type RequestBody struct {
	Args      string `json:"args"`
	Namespace string `json:"namespace"`
	Name      string `json:"name"`
	Image     string `json:"image"`
}
