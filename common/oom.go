package common

import (
	"fmt"
	"time"

	"pal-server-helper/pal"

	"github.com/shirou/gopsutil/mem"
)

type MemoryStatus struct {
	Total uint64 `json:"total"`
	Used  uint64 `json:"used"`
}

func MonitorMemoryUsage(threshold float64, checkInterval int) error {
	ticker := time.NewTicker(time.Duration(checkInterval) * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			memory, err := mem.VirtualMemory()
			if err != nil {
				fmt.Println("Failed to get memory info:", err)
				continue
			}

			usedPercent := memory.UsedPercent
			// fmt.Println(fmt.Sprintf("Current memory usage: %.2f%%", usedPercent))
			if usedPercent > threshold {
				pal.Reboot(true)
			}
		}
	}
}

func GetMemoryStats() (MemoryStatus, error) {
	memory, err := mem.VirtualMemory()
	if err != nil {
		return MemoryStatus{}, fmt.Errorf("Failed to get memory info:%v", err)
	}

	return MemoryStatus{
		Total: memory.Total,
		Used:  memory.Used,
	}, nil
}
