package systemusage

import (
	"log"
	"os/exec"
	"strconv"
	"strings"

	"github.com/yusufpapurcu/wmi"
)

type CPUUsage struct {
	Name                 string
	PercentProcessorTime uint64
}

type Win32_Process struct {
	Name           string
	WorkingSetSize uint64 
}

type DiskUsage struct {
	Name           string
	DiskBytesPersec uint64
}

func GetCPUUsage() (map[string]float64, error) {
	out, err := exec.Command("wmic", "path", "Win32_PerfFormattedData_PerfProc_Process", "get", "Name,PercentProcessorTime").Output()
	if err != nil {
		log.Printf("Failed to query CPU usage: %v", err)
		return nil, err
	}

	lines := strings.Split(string(out), "\n")
	usage := make(map[string]float64)

	for _, line := range lines[1:] {
		fields := strings.Fields(line)
		if len(fields) == 2 {
			cpu, err := strconv.ParseFloat(fields[1], 64)
			if err == nil {
				usage[fields[0]] = cpu
			}
		}
	}
	return usage, nil
}

func GetMemoryUsage() ([]Win32_Process, error) {
	var dst []Win32_Process
	query := wmi.CreateQuery(&dst, "")
	err := wmi.Query(query, &dst)
	if err != nil {
		log.Printf("Failed to query process usage: %v", err)
		return nil, err
	}
	return dst, nil
}

func GetDiskUsage() (map[string]float64, error) {
	out, err := exec.Command("wmic", "path", "Win32_PerfFormattedData_PerfDisk_LogicalDisk", "get", "Name,DiskBytesPersec").Output()
	if err != nil {
		log.Printf("Failed to query Disk usage: %v", err)
		return nil, err
	}

	lines := strings.Split(string(out), "\n")
	usage := make(map[string]float64)

	for _, line := range lines[1:] {
		fields := strings.Fields(line)
		if len(fields) == 2 {
			disk, err := strconv.ParseFloat(fields[1], 64)
			if err == nil {
				usage[fields[0]] = disk
			}
		}
	}
	return usage, nil
}
