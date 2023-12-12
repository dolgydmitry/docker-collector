package dockercl

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCpuUsagePecent(t *testing.T) {
	inData := CpuUsagePercentParams{
		ContCpuTotal:    2.359968e+09,
		SysCpu:          1.48525322e+15,
		PreContCpuTotal: 2.35901e+09,
		PreSysCpu:       1.48524918e+15,
		CpuCount:        4,
	}
	res := CpuUsagePercent(&inData)
	require.Equal(t, 0.09485148514851485, res)
}
