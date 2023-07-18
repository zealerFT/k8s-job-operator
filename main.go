package main

import (
	"flag"
	"fmt"
	"os"

	"k8s-job-operator/cmd"

	"github.com/rs/zerolog/log"
)

func main() {
	flag.Parse()
	log.Info().Msg("k8s-job-operator: is starting")
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	log.Info().Msg("k8s-job-operator: is end")

}
