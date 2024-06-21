package system

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"
)

type SystemInfo struct {
	Memory        *mem.VirtualMemoryStat `json:"memory"`
	CPUInfo       []cpu.InfoStat         `json:"cpuInfo"`
	CPUUsageRates []float64              `json:"cpuUsageRates"`
	DiskUsage     *disk.UsageStat        `json:"diskUsage"`
}

func (s SystemInfo) Print() {
	fmt.Printf("总内存: %v MB\n", s.Memory.Total/1024/1024)
	fmt.Printf("可用内存: %v MB\n", s.Memory.Available/1024/1024)
	fmt.Printf("已用内存: %v MB\n", s.Memory.Used/1024/1024)
	fmt.Printf("内存使用率: %f%%\n", s.Memory.UsedPercent)

	for _, cpuInfo := range s.CPUInfo {
		fmt.Printf("CPU型号: %s 核心数:%d\n", cpuInfo.ModelName, cpuInfo.Cores)
	}
	for i, cpuPercent := range s.CPUUsageRates {
		fmt.Printf("CPU(%d) 使用率: %v%%\n", i, cpuPercent)
	}

	fmt.Printf("总磁盘空间: %v GB\n", s.DiskUsage.Total/1024/1024/1024)
	fmt.Printf("可用磁盘空间: %v GB\n", s.DiskUsage.Free/1024/1024/1024)
	fmt.Printf("已用磁盘空间: %v GB\n", s.DiskUsage.Used/1024/1024/1024)
	fmt.Printf("磁盘使用率: %f%%\n", s.DiskUsage.UsedPercent)
}

func GetSystemUsage() (*SystemInfo, error) {
	systemInfo := &SystemInfo{}
	vm, err := mem.VirtualMemory()
	if err != nil {
		fmt.Println("获取内存使用情况失败:", err)
		return nil, err
	}
	systemInfo.Memory = vm

	// 获取 CPU 使用情况
	systemInfo.CPUUsageRates, err = cpu.Percent(time.Second, false)
	if err != nil {
		fmt.Println("获取 CPU 使用情况失败:", err)
		return nil, err
	}

	systemInfo.CPUInfo, err = cpu.Info()
	if err != nil {
		fmt.Println("获取 CPU 信息失败:", err)
		return nil, err
	}

	// 获取磁盘使用情况
	systemInfo.DiskUsage, err = disk.Usage("/")
	if err != nil {
		fmt.Println("获取磁盘使用情况失败:", err)
		return nil, err
	}
	return systemInfo, nil
}
