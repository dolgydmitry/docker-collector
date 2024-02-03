package dockercl

import (
	"context"
	// "docker-collector/pkg/collectors"
	"docker-collector/pkg/metrics"
	"encoding/json"
	"strings"
	"sync"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
	"github.com/prometheus/client_golang/prometheus"
)

var delay = 5 * time.Second

type cpu struct {
	Usage          cpuUsage `json:"cpu_usage"`
	SystemCpuUsage float64  `json:"system_cpu_usage"`
	CpuCount       int      `json:"online_cpus"`
}

type cpuUsage struct {
	Total float64 `json:"total_usage"`
}

type memory struct {
	Usage       int64       `json:"usage"`
	Limit       int64       `json:"limit"`
	MemoryStats memoryStats `json:"stats"`
}

type memoryStats struct {
	Cache int64 `json:"cache"`
}

type DockerStats struct {
	Cpu    cpu    `json:"cpu_stats"`
	PreCpu cpu    `json:"precpu_stats"`
	Memory memory `json:"memory_stats"`
}

type CnNameID struct {
	CnName   string
	CnId     string
	Stats    types.ContainerStats
	Data     DockerStats
	CpuUsage float64
	Memory   int64
}

var getStats = func(ctx context.Context, cnStuff <-chan CnNameID, cli *client.Client) <-chan CnNameID {
	/*
		Generator use container ID grep docker stats
	*/
	statsStream := make(chan CnNameID)
	go func() {
		defer close(statsStream)
		for cn := range cnStuff {
			cnData, _ := cli.ContainerStats(ctx, cn.CnId, false)
			result := DockerStats{}
			json.NewDecoder(cnData.Body).Decode(&result)
			cn.Data = result
			select {
			case <-ctx.Done():
				return
			case statsStream <- cn:
			}
		}
	}()
	return statsStream
}

var computeMetric = func(ctx context.Context, cnStuff <-chan CnNameID) <-chan CnNameID {
	respStream := make(chan CnNameID)
	go func() {
		defer close(respStream)
		for cn := range cnStuff {
			cpuPrecent := CpuUsagePercent(&CpuUsagePercentParams{
				ContCpuTotal:    cn.Data.Cpu.Usage.Total,
				SysCpu:          cn.Data.Cpu.SystemCpuUsage,
				PreContCpuTotal: cn.Data.PreCpu.Usage.Total,
				PreSysCpu:       cn.Data.PreCpu.SystemCpuUsage,
				CpuCount:        cn.Data.Cpu.CpuCount,
			})
			memory := cn.Data.Memory.Usage - cn.Data.Memory.MemoryStats.Cache
			cn.CpuUsage = cpuPrecent
			cn.Memory = memory
			select {
			case <-ctx.Done():
				return
			case respStream <- cn:
			}
		}
	}()
	return respStream
}

func updateMetric(ctx context.Context, inStream <-chan CnNameID, metricsMap map[string]map[string]prometheus.Gauge) {
	var wg sync.WaitGroup
	for t := range inStream {
		wg.Add(1)
		go func() {
			defer wg.Done()
			select {
			case <-ctx.Done():
				return
			default:
				// log.Printf("add value for container: %s", t.CnName)
				metricsMap[t.CnName][metrics.CpuName].Set(t.CpuUsage)
				metricsMap[t.CnName][metrics.MemoryName].Set(float64(t.Memory))
			}
		}()
	}
	wg.Wait()
}

func BranchPipe(ctx context.Context, channels ...<-chan CnNameID) <-chan CnNameID {
	/*
		func receive multiple channels, run in and wait until complete and multiplex it in one channel to out
	*/
	var wg sync.WaitGroup
	cnMuxStream := make(chan CnNameID)

	// worker defeniton
	cnMux := func(c <-chan CnNameID) {
		defer wg.Done()
		for i := range c {
			select {
			case <-ctx.Done():
				return
			case cnMuxStream <- i:
			}
		}
	}

	//run workers
	wg.Add(len(channels))
	for _, c := range channels {
		go cnMux(c)
	}

	// wait unitl complete
	go func() {
		wg.Wait()
		close(cnMuxStream)
	}()
	return cnMuxStream
}

var containersGen = func(ctx context.Context, containers []types.Container) <-chan CnNameID {
	outStream := make(chan CnNameID)
	go func() {
		defer close(outStream)
		for _, value := range containers {
			dockerCNName := strings.Replace(value.Names[0], "/", "", 1)
			select {
			case <-ctx.Done():
				log.Printf("close simpleGen by reason: %s" + ctx.Err().Error())
				return

			case outStream <- CnNameID{CnName: dockerCNName, CnId: value.ID}:
			}
		}

	}()
	return outStream
}

func metricProccesor(cnNameList *[]string, metricsMap map[string]map[string]prometheus.Gauge, cli *client.Client) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	filtersArgs := filters.NewArgs()
	for _, cnName := range *cnNameList {
		filtersArgs.Add("name", cnName)
	}
	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{
		Filters: filtersArgs,
	})
	if err != nil {
		log.Panic().Err(err)
	}

	updateMetric(ctx, computeMetric(ctx, getStats(ctx, containersGen(ctx, containers), cli)), metricsMap)
}

func GetAskDocker(chNameList *[]string, metricsMap map[string]map[string]prometheus.Gauge, cli *client.Client) {
	for {
		metricProccesor(chNameList, metricsMap, cli)

		// debug part
		// var m runtime.MemStats
		// runtime.ReadMemStats(&m)
		// fmt.Printf("Alloc = %v kB", m.Alloc/1024)
		// fmt.Printf("\tHeaplAlloc = %v kB", m.HeapAlloc/1024)
		// fmt.Printf("\tSys = %v MiB", m.Sys/1024/1024)
		// fmt.Printf("\tHeapObjects = %v", m.HeapObjects)
		// fmt.Printf("\tNumGC = %v\n", m.NumGC)

		time.Sleep(5 * time.Second)
	}
}
