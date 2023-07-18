package config

import (
	"fmt"
	"os"

	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v3"
)

var (
	options *AppConfig
)

func init() {
	cfg := &AppConfig{}
	var err error
	var file []byte
	file, err = os.ReadFile("config/app.test.yaml")
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(file, cfg)
	if err != nil {
		panic(err)
	}
	if err != nil {
		log.Panic().Msgf("%+v", err)
	}

	fmt.Println("-------------------------------------")

	options = cfg
}

// AppConfig store all configuration options
type AppConfig struct {
	AppName        string `yaml:"app_name" env:"APP_NAME" json:"app_name,omitempty"`
	AppEnv         string `yaml:"app_environment" env:"APP_ENVIRONMENT" json:"app_env,omitempty"`
	Debug          bool   `yaml:"debug" env:"APP_DEBUG" json:"debug,omitempty"`
	AppRole        string `yaml:"role" env:"APP_ROLE" json:"app_role,omitempty"`
	HTTPServerAddr string `yaml:"http_server_addr" env:"HTTP_SERVER_ADDR" json:"http_server_addr,omitempty"`

	KubeSetting KubeSetting `json:"kubeSetting" yaml:"kube_setting"`
}

type KubeSetting struct {
	KubeconfigFile string `yaml:"kubeconfig_file"`
}

func SetRole(role string) {
	options.AppRole = role
}

// Options return application config options
func Options() *AppConfig {
	return options
}
