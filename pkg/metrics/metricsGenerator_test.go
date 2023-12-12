package metrics

import (
	"testing"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/stretchr/testify/require"
)

func TestMetricsGen(t *testing.T) {
	cns := []string{"front-cn", "backend-cn", "db-cn"}
	metrics := MetricsGen(cns)
	require.NotEmpty(t, metrics)
	require.Len(t, metrics, 3)
	for _, cn := range cns {
		require.NotEmpty(t, metrics[cn])
		require.Len(t, metrics[cn], 2)
		for _, val := range metrics[cn] {
			_, ok := val.(prometheus.Gauge)
			require.True(t, ok)
		}
	}

}
