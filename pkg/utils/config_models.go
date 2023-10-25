package utils

// type ConfigMetric struct {
// 	Name      string
// 	Type      string `json:"type"`
// 	ValueType string `json:"value_type"`
// 	Help      string `json:"help"`
// 	MetricObj interface{}
// }

// type Container struct {
// 	ChName  string          `json:"contianer_name"`
// 	CnID    string          `json:"contianer_id"`
// 	Metrics []*ConfigMetric `json:"metrics"`
// }

// type Config struct {
// 	Containers []*Container `json:"containers"`
// }

// type ConfigMetric struct {
// 	Name      string
// 	Type      string `mapstructure:"type"`
// 	ValueType string `mapstructure:"value_type"`
// 	Help      string `mapstructure:"help"`
// 	MetricObj prometheus.Counter
// }

// type Container struct {
// 	ChName  string          `mapstructure:"contianer_name"`
// 	CnID    string          `mapstructure:"contianer_id"`
// 	Metrics []*ConfigMetric `mapstructure:"metrics"`
// }

// type Config struct {
// 	Containers []*Container `mapstructure:"containers"`
// }

type Config struct {
	Containers    []string `mapstructure:"containers"`
	ServerAddress string   `mapstructure:"server_address"`
}
