package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"

	"github.com/rs/zerolog/log"

	"github.com/spf13/viper"
)

var (
	_, b, _, _ = runtime.Caller(0)
	Root       = filepath.Join(filepath.Dir(b), "../..")
)

const (
	// filePath = "config.json"
	filePath = "config.json"
)

func LoadInitialConfig() *Config {
	loadPath := fmt.Sprintf("%v/%v", Root, filePath)
	config := &Config{}
	file, err := os.Open(loadPath)
	if err != nil {
		log.Panic().Err(err)
	}
	defer file.Close()
	byteValue, err := io.ReadAll(file)
	if err != nil {
		log.Panic().Err(err)
	}
	json.Unmarshal(byteValue, config)
	return config
}

func LoadInitialConfigYaml() (*Config, error) {
	config := &Config{}
	viper.AddConfigPath(Root)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	viper.Unmarshal(&config)
	return config, nil
}

// func CreateMetricsConfig() map[string][]*metrics.Metric {
// 	result := make(map[string][]*metrics.Metric)
// 	builder := &metrics.MetricBuilder{}
// 	config := LoadInitialConfig()
// 	for _, value := range config.CnNames {
// 		var metricsList []*metrics.Metric
// 		for _, metricValue := range config.MetricsValues[value] {
// 			newMetric := builder.AddNewMetric(&metrics.MetricParam{
// 				Name:  value,
// 				Type:  "counter",
// 				Value: metricValue,
// 			})
// 			metricsList = append(metricsList, newMetric)
// 		}
// 		result[value] = metricsList
// 	}
// 	return result
// }

// func CreateMetricsConfig() []*MetricsGroup {
// 	var result []*MetricsGroup
// 	builder := &metrics.MetricBuilder{}
// 	config := LoadInitialConfig()
// 	for _, value := range config.CnNames {
// 		var metricsList []*metrics.Metric
// 		for _, metricValue := range config.MetricsValues[value] {
// 			newMetric := builder.AddNewMetric(&metrics.MetricParam{
// 				Name:  value,
// 				Type:  "counter",
// 				Value: metricValue,
// 			})
// 			metricsList = append(metricsList, newMetric)
// 		}
// 		metricGroup := &MetricsGroup{
// 			GroupName:   value,
// 			MetricsList: metricsList,
// 		}
// 		result = append(result, metricGroup)
// 	}
// 	return result
// }
