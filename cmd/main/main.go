package main

import (
	// "docker-collector/pkg/server"

	"docker-collector/pkg/dockercl"
	"docker-collector/pkg/metrics"
	"docker-collector/pkg/server"
	"docker-collector/pkg/utils"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	// "net/http"
	"runtime/debug"

	// _ "net/http/pprof"

	"github.com/common-nighthawk/go-figure"
	"github.com/docker/docker/client"
)

func main() {

	// debug pprof part
	// go func() {
	// 	log.Println(http.ListenAndServe("localhost:6060", nil))
	// }()

	myFigure := figure.NewColorFigure("Docker collector", "", "green", true)
	myFigure.Print()

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	debug.SetMemoryLimit(14096000)
	config, err := utils.LoadInitialConfigYaml()
	if err != nil {
		log.Fatal().Msgf("cannot load config file: %s", err)
	}

	var cli, _ = client.NewClientWithOpts(client.FromEnv)
	defer cli.Close()

	metrics := metrics.MetricsGen(config.Containers)
	go dockercl.GetAskDocker(&config.Containers, metrics, cli)

	server.RunServerApp(config.ServerAddress)
}
