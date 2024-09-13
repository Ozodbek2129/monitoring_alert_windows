package gpu

import (
	"log"

	"github.com/yusufpapurcu/wmi"
)

type Win32_VideoController struct {
	Name                        string
	AdapterRAM                  uint64
	CurrentVerticalResolution   uint32
	CurrentHorizontalResolution uint32
	LoadPercentage              uint16
}

func GetGPUUsage() (uint16, error) {
	var dst []Win32_VideoController
	query := wmi.CreateQuery(&dst, "")

	err := wmi.Query(query, &dst)
	if err != nil {
		log.Printf("Failed to query GPU usage: %v", err)
		return 0, err
	}

	if len(dst) > 0 {
		return dst[0].LoadPercentage, nil
	}

	return 0, nil
}
