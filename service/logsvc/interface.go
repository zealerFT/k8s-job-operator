package logsvc

import (
	"k8s-job-operator/service/snowflakesvc"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func NewLogger(snowflake *snowflakesvc.Snowflake) zerolog.Logger {
	return log.With().Str("request_id", snowflake.Generate().String()).Logger()
}
