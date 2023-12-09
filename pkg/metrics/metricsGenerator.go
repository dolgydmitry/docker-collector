package metrics

import (
	"fmt"
	"log"

	"strings"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var metricList = []string{"cpu", "memory"}
var CpuName = metricList[0]
var MemoryName = metricList[1]

func createCounterV1(cnName string) map[string]prometheus.Counter {
	var result = map[string]prometheus.Counter{}
	for _, metricNameIn := range metricList {
		metricNameSands := fmt.Sprintf("%v_%v_%v", appName, cnName, metricNameIn)
		metricName := strings.Replace(metricNameSands, "-", "_", -1)
		log.Printf("Create metric : %s", cnName)
		newMetric := promauto.NewCounter(prometheus.CounterOpts{
			Name: metricName,
			Help: help,
		})
		result[metricNameIn] = newMetric
	}
	return result
}

func createGaugerV1(cnName string) map[string]prometheus.Gauge {
	var result = map[string]prometheus.Gauge{}
	for _, metricNameIn := range metricList {
		metricNameSands := fmt.Sprintf("%v_%v_%v", appName, cnName, metricNameIn)
		metricName := strings.Replace(metricNameSands, "-", "_", -1)
		log.Printf("Create metric: %s", metricNameSands)
		newMetric := promauto.NewGauge(prometheus.GaugeOpts{
			Name: metricName,
			Help: help,
		})
		result[metricNameIn] = newMetric
	}
	return result
}

func MetricsGen(cnList []string) map[string]map[string]prometheus.Gauge {
	var res = map[string]map[string]prometheus.Gauge{}
	for _, cn := range cnList {
		res[cn] = createGaugerV1(cn)
	}
	return res
}
