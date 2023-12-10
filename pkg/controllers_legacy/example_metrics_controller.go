package controllers

import (
	"docker-collector/pkg/metrics"

	"github.com/prometheus/client_golang/prometheus"
)

func GetRegistry() *prometheus.Registry {
	reg := prometheus.NewRegistry()
	builder := &metrics.BuilderMetrics{}
	m := builder.NewMetrics(reg)
	m.CpuUsage.Set(98)
	m.MemoryUsage.Set(77)
	return reg

}
