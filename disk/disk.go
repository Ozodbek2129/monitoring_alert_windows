package disk

import (
	"github.com/shirou/gopsutil/disk"
)

func GetDiskUsage() (float64, error) {
	usage, err := disk.Usage("/")
	if err != nil {
		return 0, err
	}
	return usage.UsedPercent, nil
}
