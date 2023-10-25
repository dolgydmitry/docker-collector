package main

import (
	// "docker-collector/pkg/server"

	"docker-collector/pkg/collectors"
	"docker-collector/pkg/controllers"
	"docker-collector/pkg/metrics"
	"docker-collector/pkg/server"
	"docker-collector/pkg/utils"
	"log"
	"time"

	"github.com/common-nighthawk/go-figure"
)

const (
	host = ":8098"
)

func main() {
	myFigure := figure.NewColorFigure("Docker collector", "", "green", true)
	myFigure.Print()

	config, err := utils.LoadInitialConfigYaml()
	if err != nil {
		log.Fatal("cannot load config file: ", err)
	}
	ticker := time.NewTicker(5 * time.Second)
	quit := make(chan struct{})

	// initilize metrics
	log.Print("Start metrics creating")
	builder := &metrics.MetricBuilder{}
	builder.Constructor()
	builder.CreateMetric(&config.Containers)
	metrics := builder.Metrics
	log.Print("Finish metrics creating")

	// collectors list
	var colList []collectors.CollectorExp

	// create docker collector
	DockerCli := &collectors.DockerCli{CnObserve: metrics}
	DockerCli.Constructor()
	colList = append(colList, DockerCli)

	// run all colletor
	log.Print("Run docker collector")
	go controllers.CollectMetric(
		controllers.MainControllerParams{
			Col:    colList,
			Ticker: ticker,
			Quit:   quit,
		},
	)

	server.RunServerApp(config.ServerAddress)
}
