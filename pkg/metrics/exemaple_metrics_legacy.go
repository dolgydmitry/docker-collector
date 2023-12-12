package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

type BaseMetrics struct {
	CpuUsage    prometheus.Gauge
	MemoryUsage prometheus.Gauge
}

type BuilderMetrics struct{}

func (b *BuilderMetrics) NewMetrics(reg prometheus.Registerer) *BaseMetrics {
	m := &BaseMetrics{
		CpuUsage: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "cpu_usage",
			Help: "Current cpu usage",
		}),
		MemoryUsage: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "memory_usage",
			Help: "Current memory usage",
		}),
	}
	reg.MustRegister(m.CpuUsage)
	reg.MustRegister(m.MemoryUsage)
	return m
}
