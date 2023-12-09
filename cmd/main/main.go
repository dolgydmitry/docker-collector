package main

import (
	// "docker-collector/pkg/server"

	"docker-collector/pkg/dockercl"
	"docker-collector/pkg/metrics"
	"docker-collector/pkg/server"
	"docker-collector/pkg/utils"
	"log"
	"net/http"
	"runtime/debug"

	_ "net/http/pprof"

	"github.com/common-nighthawk/go-figure"
	"github.com/docker/docker/client"
)

const (
	host = ":8098"
)

// var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to `file`")
// var memprofile = flag.String("memprofile", "", "write memory profile to `file`")
// var memprofile = flag.String("memprofile", "", "write memory profile to this file")

func main() {

	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()
	// flag.Parse()
	// if *cpuprofile != "" {
	// 	f, err := os.Create(*cpuprofile)
	// 	if err != nil {
	// 		log.Fatal("could not create CPU profile: ", err)
	// 	}
	// 	defer f.Close() // error handling omitted for example
	// 	if err := pprof.StartCPUProfile(f); err != nil {
	// 		log.Fatal("could not start CPU profile: ", err)
	// 	}
	// 	defer pprof.StopCPUProfile()
	// }

	myFigure := figure.NewColorFigure("Docker collector", "", "green", true)
	myFigure.Print()

	// config, err := utils.LoadInitialConfigYaml()
	// if err != nil {
	// 	log.Fatal("cannot load config file: ", err)
	// }
	// reg := prometheus.NewRegistry()
	// metrics := collectors.CreateMetrics(reg, config.Containers)

	// collectors.AddContainerID(config.Containers, metrics)
	// var mem runtime.MemStats
	// go func() {
	// 	for {
	// 		runtime.ReadMemStats(&mem)
	// 		log.Printf("alloc [%v] \t heapAlloc [%v] \n", mem.Alloc, mem.HeapAlloc)
	// 		time.Sleep(5 * time.Second)
	// 		collectors.SetMetrics(config.Containers, metrics)
	// 	}
	// }()

	// Expose the registered metrics via HTTP.
	// http.Handle("/metrics", promhttp.HandlerFor(
	// 	reg,
	// 	promhttp.HandlerOpts{
	// 		// Opt into OpenMetrics to support exemplars.
	// 		EnableOpenMetrics: true,
	// 		// Pass custom registry
	// 		Registry: reg,
	// 	},
	// ))
	// log.Fatal(http.ListenAndServe("0.0.0.0:8099", nil))

	//------------------------------------------------
	// ticker := time.NewTicker(5 * time.Second)
	// quit := make(chan struct{})

	// initilize metrics
	// log.Print("Start metrics creating")
	// builder := &metrics.MetricBuilder{}
	// builder.Constructor()
	// builder.CreateMetric(&config.Containers)
	// metrics := builder.Metrics
	// log.Print("Finish metrics creating")

	//-------------------------------------------------------
	// collectors list
	// var colList []collectors.CollectorExp

	// create docker collector
	// DockerCli := collectors.DockerCli{CnObserve: metrics}
	// DockerCli.Constructor()
	// // colList = append(colList, DockerCli)

	// DockerCli.GetMetricsValue()
	// // run all colletor
	// log.Print("Run docker collector")

	// controllers.CollectMetric(
	// 	// controllers.MainControllerParams{
	// 	// 	Col:    colList,
	// 	// 	Ticker: ticker,
	// 	// 	Quit:   quit,
	// 	// },
	// 	&DockerCli,
	// )

	// debug.SetMemoryLimit(4096000)
	// debug.SetGCPercent(100)
	// stats := debug.GCStats{}
	// debug.ReadGCStats(&stats)
	// fmt.Println(stats)

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
	//FindHavlakLoops(cfgraph, lsgraph)
	// if *memprofile != "" {
	// 	f, err := os.Create(*memprofile)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	pprof.WriteHeapProfile(f)
	// 	f.Close()
	// 	return
	// }
}
