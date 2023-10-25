package controllers

import (
	"docker-collector/pkg/collectors"
	"time"
)

type MainControllerParams struct {
	Col    []collectors.CollectorExp
	Ticker *time.Ticker
	Quit   chan struct{}
}

func CollectMetric(param MainControllerParams) {
	for {
		select {
		case <-param.Ticker.C:
			for _, col := range param.Col {
				go col.GetMetricsValue()

			}
		case <-param.Quit:
			param.Ticker.Stop()
			return
		}
	}
}
