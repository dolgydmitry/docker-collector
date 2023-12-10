package collectors

import (
	"context"
	"docker-collector/pkg/metrics"
	"log"
	"reflect"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
)

type DockerCli struct {
	CnObserve   map[string]metrics.CnMetrics
	CnsID       map[string]string
	cli         *client.Client
	ctx         context.Context
	filtersArgs filters.Args
}

func (d *DockerCli) Constructor() {
	var err error
	d.cli, err = client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		log.Panic(err)
	}
	d.ctx = context.Background()
	d.CnsID = map[string]string{}
	d.filtersArgs = filters.NewArgs()
	for cnName, _ := range d.CnObserve {
		d.filtersArgs.Add("name", cnName)
	}
	d.addContainerID()
}

func (d *DockerCli) addContainerID() {
	containers, err := d.cli.ContainerList(d.ctx, types.ContainerListOptions{
		Filters: d.filtersArgs,
	})
	if err != nil {
		log.Panic(err)
	}
	for _, value := range containers {
		dockerCNName := strings.Replace(value.Names[0], "/", "", 1)
		d.CnsID[dockerCNName] = value.ID
	}
	// Check if container didn't find by name
	CnsIDKeys := reflect.ValueOf(d.CnsID).MapKeys()
	CnObserveKeys := reflect.ValueOf(d.CnObserve).MapKeys()
	if !reflect.DeepEqual(CnsIDKeys, CnObserveKeys) {
		for key, _ := range d.CnObserve {
			_, ok := d.CnsID[key]
			if !ok {
				log.Printf("WARN container: %s didn't found in the docker daemon's host", key)
			}
		}
	}
}

type DockerStats struct {
	Cpu    cpu    `json:"cpu_stats"`
	PreCpu cpu    `json:"precpu_stats"`
	Memory memory `json:"memory_stats"`
}

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

func (d *DockerCli) getOneMetricValue(cnName, cnId string) {
	// log.Printf("get metric to container %s", cnName)
	// value, err := d.cli.ContainerStats(d.ctx, cnId, false)
	// if err != nil {
	// 	log.Printf("cannot get container stats for container: %s, by error: %v", cnName, err)
	// } else {
	// 	result := DockerStats{}
	// 	json.NewDecoder(value.Body).Decode(&result)

	// 	cpuPrecent := cpuUsagePercent(&CpuUsagePercentParams{
	// 		contCpuTotal:    result.Cpu.Usage.Total,
	// 		sysCpu:          result.Cpu.SystemCpuUsage,
	// 		preContCpuTotal: result.PreCpu.Usage.Total,
	// 		preSysCpu:       result.PreCpu.SystemCpuUsage,
	// 		cpuCount:        result.Cpu.CpuCount,
	// 	})
	// 	memory := result.Memory.Usage - result.Memory.MemoryStats.Cache
	// 	// test1 := *d.CnObserve[cnName].Gauges["cpu"]
	// 	// test1.Set(cpuPrecent)

	// 	// test2 := *d.CnObserve[cnName].Gauges["memory"]
	// 	// test2.Set(float64(memory))
	// 	d.CnObserve[cnName].Gauges["cpu"].Set(cpuPrecent)
	// 	d.CnObserve[cnName].Gauges["memory"].Set(float64(memory))
	// }
	d.CnObserve[cnName].Gauges["cpu"].Set(10)
	d.CnObserve[cnName].Gauges["memory"].Set(float64(20))
}

func (d *DockerCli) GetMetricsValue() {
	d.addContainerID()
	for key, value := range d.CnsID {
		d.getOneMetricValue(key, value)
	}
}
