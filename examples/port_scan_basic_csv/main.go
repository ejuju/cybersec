package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/ejuju/cybersec/pkg/netutil"
)

func main() {
	// store start timestamp for stats
	start := time.Now()

	// scan ports
	ports := netutil.PortRange(0, 65535)
	networks := map[string]struct{}{"tcp4": {}, "tcp6": {}, "udp4": {}, "udp6": {}}
	scanResults := netutil.ScanPorts(netutil.PortScanConfiguration{
		Ports:    ports,
		Networks: networks,
		Host:     "127.0.0.1",
		Scanner: &netutil.DefaultPortScanner{
			ReadTimeout:      50 * time.Millisecond,
			DialTimeout:      500 * time.Millisecond,
			BannerBufferSize: 1024,
		},
	})

	// show results stats
	timeToExecute := time.Now().Sub(start)
	fmt.Printf("Scanned %d ports on %d networks in %s (%.2f ports/sec)\n",
		len(ports),
		len(networks),
		timeToExecute,
		float64(len(ports))/timeToExecute.Seconds(),
	)

	// create csv file for open ports
	f, err := os.Create("open_ports.csv")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	csvWriter := csv.NewWriter(f)
	err = csvWriter.Write([]string{"Network", "Port", "IsOpen", "Banner", "Info"})

	// write open ports to csv file
	for _, scanResult := range scanResults {
		if scanResult.IsOpen {
			err = csvWriter.Write(scanResultToCSVRecord(scanResult))
			if err != nil {
				panic(err)
			}
		}
	}
	csvWriter.Flush()

	err = csvWriter.Error()
	if err != nil {
		panic(err)
	}
}

// utility function to transform a port scan result to a CSV-writable record
func scanResultToCSVRecord(result *netutil.PortScanResult) []string {
	return []string{
		result.Network,
		strconv.Itoa(result.Port),
		strconv.FormatBool(result.IsOpen),
		strings.TrimSuffix(result.Banner, "\r\n"),
		result.Info,
	}
}
