package metrics

import (
	"fmt"
	"log"

	"strings"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

const (
	appName    = "docker-collector"
	metricType = "counter"
	help       = "custom stubs"
)

// var valueTypes = []string{"cpu", "memory"}

var valueTypes = map[string][]string{
	"counter": []string{},
	"gauge": []string{
		"cpu",
		"memory",
	},
}

type CnMetrics struct {
	Gauges   map[string]prometheus.Gauge
	Counters map[string]prometheus.Counter
}

type MetricBuilder struct {
	ValueTypes map[string][]string
	Metrics    map[string]CnMetrics
}

func (builder *MetricBuilder) Constructor() {
	builder.ValueTypes = valueTypes
	builder.Metrics = make(map[string]CnMetrics)
}

func (builder *MetricBuilder) CreateMetric(cnNames *[]string) {
	for _, cnName := range *cnNames {
		log.Printf("Create metric: %s", cnName)
		builder.metricsByContainer(cnName)
	}
}

func (builder *MetricBuilder) metricsByContainer(cnName string) {
	metrics := CnMetrics{}
	for metricType, metricNames := range builder.ValueTypes {
		if len(metricNames) != 0 {
			switch metricType {
			case "counter":
				metrics.Counters = createCounter(cnName, metricNames)
			case "gauge":
				metrics.Gauges = createGauger(cnName, metricNames)
			}
		}
	}
	builder.Metrics[cnName] = metrics
}

func createCounter(cnName string, metricNames []string) map[string]prometheus.Counter {
	var result = map[string]prometheus.Counter{}
	for _, metricNameIn := range metricNames {
		metricNameSands := fmt.Sprintf("%v_%v_%v", appName, cnName, metricNameIn)
		metricName := strings.Replace(metricNameSands, "-", "_", -1)
		log.Printf("Create metric: %s", cnName)
		newMetric := promauto.NewCounter(prometheus.CounterOpts{
			Name: metricName,
			Help: help,
		})
		result[metricNameIn] = newMetric
	}
	return result
}

func createGauger(cnName string, metricNames []string) map[string]prometheus.Gauge {
	var result = map[string]prometheus.Gauge{}
	for _, metricNameIn := range metricNames {
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

// func (builder *MetricBuilderCounter) CreateMetric(cnNames *[]string) map[string]map[string]prometheus.Counter {
// 	result := map[string]map[string]prometheus.Counter{}
// 	for _, cnName := range *cnNames {
// 		metrics := map[string]prometheus.Counter{}
// 		for _, valueType := range builder.ValueTypes {
// 			metricNameSands := fmt.Sprintf("%v_%v_%v", appName, cnName, valueType)
// 			metricName := strings.Replace(metricNameSands, "-", "_", -1)
// 			newMetric := promauto.NewCounter(prometheus.CounterOpts{
// 				Name: metricName,
// 				Help: help,
// 			})
// 			metrics[valueType] = newMetric
// 		}
// 		result[cnName] = metrics
// 	}
// 	return result
// }
