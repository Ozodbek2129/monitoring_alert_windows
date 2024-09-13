package cpu

import (
	"github.com/shirou/gopsutil/cpu"
)

func GetCPUUsage() (float64, error) {
	percent, err := cpu.Percent(0, false)
	if err != nil {
		return 0, err
	}
	return percent[0], nil
}
