package monitor

import (
	"fmt"
	"runtime"
	"time"

	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/disk"
	"github.com/shirou/gopsutil/v4/mem"
)

type SystemMetrics struct {
	CpuCores int     `json:"cpuCores"`
	CpuUsage float64 `json:"cpuUsage,string"`
	CpuInfo  string  `json:"cpuInfo"`

	MemTotal        uint64  `json:"memTotal,string"`
	MemUsage        uint64  `json:"memUsage,string"`
	MemFree         uint64  `json:"memFree,string"`
	MemUsagePercent float64 `json:"memUsagePercent,string"`

	DiskTotal        uint64  `json:"diskTotal,string"`
	DiskUsage        uint64  `json:"diskUsage,string"`
	DiskFree         uint64  `json:"diskFree,string"`
	DiskUsagePercent float64 `json:"diskUsagePercent,string"`

	Goroutines int    `json:"goroutines,string"`
	HeapAlloc  uint64 `json:"heapAlloc,string"`
	HeapSys    uint64 `json:"heapSys,string"`
	LastGC     string `json:"lastGC"`
	NextGC     string `json:"nextGC"`
	NumGC      uint32 `json:"numGC,string"`
}

func GatherSystemMetrics() (*SystemMetrics, error) {
	result := &SystemMetrics{}
	memory, err := mem.VirtualMemory()
	if err != nil {
		return nil, err
	}
	result.MemTotal = memory.Total / 1000_000
	result.MemUsage = memory.Used / 1000_000
	result.MemFree = memory.Free / 1000_000
	result.MemUsagePercent = memory.UsedPercent

	d, err := disk.Usage("/")
	if err != nil {
		return nil, err
	}
	result.DiskTotal = d.Total / 1000_000_000
	result.DiskUsage = d.Used / 1000_000_000
	result.DiskFree = d.Free / 1000_000_000
	result.DiskUsagePercent = d.UsedPercent

	cpuInfo, err := cpu.Info()
	if err != nil {
		return nil, err
	}
	result.CpuCores = runtime.NumCPU()
	cpuPercent, err := cpu.Percent(0, false)
	if err != nil {
		return nil, err
	}
	result.CpuUsage = cpuPercent[0]
	result.Goroutines = runtime.NumGoroutine()
	result.CpuInfo = fmt.Sprintf("%s %s", cpuInfo[0].ModelName, cpuInfo[0].VendorID)
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)
	result.NumGC = memStats.NumGC
	result.LastGC = time.UnixMilli(int64(memStats.LastGC / 1000000)).Format(time.DateTime)
	result.HeapAlloc = memStats.HeapAlloc / 1000_000
	result.HeapSys = memStats.HeapSys / 1000_000
	return result, nil
}
