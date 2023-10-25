package collectors

type CpuUsagePercentParams struct {
	contCpuTotal    float64
	sysCpu          float64
	preContCpuTotal float64
	preSysCpu       float64
	cpuCount        int
}

type MemoryUsageParams struct {
	usage int64
	limit int64
	cache int64
}

// Calculate CPU usage for unix
func cpuUsagePercent(params *CpuUsagePercentParams) float64 {
	cpuDelta := params.contCpuTotal - params.preContCpuTotal
	systemDelta := params.sysCpu - params.preSysCpu
	var cpuPercent = 0.0
	if systemDelta > 0.0 && cpuDelta > 0.0 {
		cpuPercent = ((cpuDelta / systemDelta) * float64(params.cpuCount)) * 100.0
	}
	return cpuPercent
}

// Calucalte memory usage for future usage
func memoryUsage(params *MemoryUsageParams) (float64, float64) {
	used_memory := params.usage - params.cache
	memory_usage_percent := (used_memory / params.limit) * 100.0
	return float64(memory_usage_percent), float64(used_memory)
}
