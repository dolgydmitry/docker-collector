package metrics

import (
	"testing"

	"docker-collector/pkg/utils"

	"github.com/stretchr/testify/require"
)

func TestMetricBuilde(t *testing.T) {
	config := utils.Config{
		Containers: []string{"backend-db", "worker-db"},
	}
	builder := &MetricBuilder{}
	builder.Constructor()
	builder.CreateMetric(&config.Containers)
	require.NotEmpty(t, builder.Metrics)
	for _, value := range builder.Metrics {
		require.NotEmpty(t, value)
		require.NotEmpty(t, value.Gauges)
		require.Len(t, value.Gauges, 2)
		require.Len(t, value.Counters, 0)
		for _, metric := range value.Gauges {
			require.NotEmpty(t, metric)
		}
	}
}
