package network

import (
	"github.com/shirou/gopsutil/net"
)

func GetNetworkStats() (uint64, error) {
	ioCounters, err := net.IOCounters(false)
	if err != nil {
		return 0, err
	}
	return ioCounters[0].BytesRecv, nil
}
