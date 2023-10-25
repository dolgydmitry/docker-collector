package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLoadInitialConfigYaml(t *testing.T) {
	config, err := LoadInitialConfigYaml()
	require.Empty(t, err)
	require.NotEmpty(t, config)
	require.NotEqual(t, 0, len(config.Containers))
	for _, cn := range config.Containers {
		require.NotEmpty(t, cn)
		require.IsType(t, "string", cn)
	}
}

// func TestLoadInitialConfigYaml(t *testing.T) {
// 	config, err := LoadInitialConfigYaml()
// 	require.Empty(t, err)
// 	require.NotEmpty(t, config)
// 	require.NotEqual(t, 0, len(config.Containers))

// 	for _, cn := range config.Containers {
// 		require.NotEmpty(t, cn)
// 		require.NotEqual(t, 0, len(cn.Metrics))
// 		for _, metric := range cn.Metrics {
// 			fmt.Println(metric)
// 			require.NotEmpty(t, metric)
// 			require.NotEmpty(t, metric.ValueType)
// 			require.NotEmpty(t, metric.Type)
// 			require.NotEmpty(t, metric.Help)
// 		}
// 	}
// }
