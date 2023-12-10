package collectors

type CpuUsagePercentParams struct {
	ContCpuTotal    float64
	SysCpu          float64
	PreContCpuTotal float64
	PreSysCpu       float64
	CpuCount        int
}

type MemoryUsageParams struct {
	Usage int64
	Limit int64
	Cache int64
}

// Calculate CPU usage for unix
func CpuUsagePercent(params *CpuUsagePercentParams) float64 {
	cpuDelta := params.ContCpuTotal - params.PreContCpuTotal
	systemDelta := params.SysCpu - params.PreSysCpu
	var cpuPercent = 0.0
	if systemDelta > 0.0 && cpuDelta > 0.0 {
		cpuPercent = ((cpuDelta / systemDelta) * float64(params.CpuCount)) * 100.0
	}
	return cpuPercent
}

// Calucalte memory usage for future usage
func MemoryUsage(params *MemoryUsageParams) (float64, float64) {
	used_memory := params.Usage - params.Cache
	memory_usage_percent := (used_memory / params.Limit) * 100.0
	return float64(memory_usage_percent), float64(used_memory)
}
