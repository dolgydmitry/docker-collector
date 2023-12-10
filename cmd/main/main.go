package main

import (
	// "docker-collector/pkg/server"

	"docker-collector/pkg/dockercl"
	"docker-collector/pkg/metrics"
	"docker-collector/pkg/server"
	"docker-collector/pkg/utils"
	"log"

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

	debug.SetMemoryLimit(14096000)
	config, err := utils.LoadInitialConfigYaml()
	if err != nil {
		log.Fatal("cannot load config file: ", err)
	}

	var cli, _ = client.NewClientWithOpts(client.FromEnv)
	defer cli.Close()

	metrics := metrics.MetricsGen(config.Containers)
	go dockercl.GetAskDocker(&config.Containers, metrics, cli)

	server.RunServerApp(config.ServerAddress)
}
