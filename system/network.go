package system

import (
	"fmt"
	"log"
	stdnet "net"
	"os"

	psnet "github.com/shirou/gopsutil/v3/net"
)

func GetNetworkSpeed(prevRecv, prevSent uint64) (uint64, uint64, uint64, uint64) {
	stats, err := psnet.IOCounters(false)
	if err != nil || len(stats) == 0 {
		return 0, 0, prevRecv, prevSent
	}
	recv := stats[0].BytesRecv
	sent := stats[0].BytesSent
	download := recv - prevRecv
	upload := sent - prevSent
	return download, upload, recv, sent
}

func FormatSpeed(bytesPerSec uint64) string {
	switch {
	case bytesPerSec < 1024:
		return fmt.Sprintf("%d B/s", bytesPerSec)
	case bytesPerSec < 1024*1024:
		return fmt.Sprintf("%d KB/s", bytesPerSec/1024)
	default:
		return fmt.Sprintf("%.1f MB/s", float64(bytesPerSec)/1024/1024)
	}
}

func GetHostname() string {
	hostname, err := os.Hostname()
	if err != nil {
		return "unknown"
	}
	return hostname
}

func GetNetworkName() string {
	stats, err := psnet.IOCounters(true)
	if err != nil || len(stats) == 0 {
		return "unknown"
	}
	for _, s := range stats {
		if s.Name != "lo" {
			return s.Name
		}
	}
	return "unknown"
}

func GetIP() string {
	conn, err := stdnet.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Println("IP error:", err)
		return "unknown"
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*stdnet.UDPAddr)
	return localAddr.IP.String()
}
