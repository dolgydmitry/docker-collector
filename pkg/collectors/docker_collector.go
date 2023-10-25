package collectors

import (
	"context"
	"docker-collector/pkg/metrics"
	"encoding/json"
	"log"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

type DockerCli struct {
	CnObserve map[string]metrics.CnMetrics
	CnsID     map[string]string
	cli       *client.Client
	ctx       context.Context
}

func (d *DockerCli) Constructor() {
	var err error
	d.cli, err = client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		log.Panic(err)
	}
	d.ctx = context.Background()
	d.CnsID = map[string]string{}
	d.addContainerID()
}

func (d *DockerCli) addContainerID() {
	containers, err := d.cli.ContainerList(d.ctx, types.ContainerListOptions{})
	if err != nil {
		log.Panic(err)
	}
	for cnName, _ := range d.CnObserve {
		for _, dockerCn := range containers {
			dockerCNName := strings.Replace(dockerCn.Names[0], "/", "", 1)
			if cnName == dockerCNName {
				// log.Printf("Update container ID for container: %s", cnName)
				d.CnsID[cnName] = dockerCn.ID
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
	value, err := d.cli.ContainerStats(d.ctx, cnId, false)
	if err != nil {
		log.Printf("cannot get container stats for container: %s, by error: %v", cnName, err)
	} else {
		result := DockerStats{}
		json.NewDecoder(value.Body).Decode(&result)
		cpuUsagePercentParams := &CpuUsagePercentParams{
			contCpuTotal:    result.Cpu.Usage.Total,
			sysCpu:          result.Cpu.SystemCpuUsage,
			preContCpuTotal: result.PreCpu.Usage.Total,
			preSysCpu:       result.PreCpu.SystemCpuUsage,
			cpuCount:        result.Cpu.CpuCount,
		}
		cpuPrecent := cpuUsagePercent(cpuUsagePercentParams)
		memory := result.Memory.Usage - result.Memory.MemoryStats.Cache
		d.CnObserve[cnName].Gauges["cpu"].Set(cpuPrecent)
		d.CnObserve[cnName].Gauges["memory"].Set(float64(memory))
	}
}

func (d *DockerCli) GetMetricsValue() {
	d.addContainerID()
	for key, value := range d.CnsID {
		go d.getOneMetricValue(key, value)
	}
}
