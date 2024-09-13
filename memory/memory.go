package memory

import (
	"log"

	"github.com/shirou/gopsutil/mem"
)

func GetMemoryUsage() (float64, error) {
	v, err := mem.VirtualMemory()
	if err != nil {
		log.Printf("Failed to get memory usage: %v", err)
		return 0, err
	}

	return v.UsedPercent, nil
}
