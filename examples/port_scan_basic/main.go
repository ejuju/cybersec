package main

import (
	"fmt"
	"time"

	"github.com/ejuju/cybersec/pkg/netutil"
)

func main() {
	// store start timestamp for stats
	start := time.Now()

	// scan ports
	ports := netutil.PortRange(0, 65535)
	networks := map[string]struct{}{"tcp4": {}, "tcp6": {}}
	scanResults := netutil.ScanPorts(netutil.PortScanConfiguration{
		Ports:    ports,
		Networks: networks,
		Host:     "127.0.0.1",
		Scanner: &netutil.DefaultPortScanner{
			ReadTimeout:      100 * time.Millisecond,
			DialTimeout:      1 * time.Second,
			BannerBufferSize: 1024,
		},
	})

	// show results stats
	timeToExecute := time.Now().Sub(start)
	fmt.Printf(
		"Scanned %d ports on %d networks in %s (%.2f ports/sec)\n",
		len(ports),
		len(networks),
		timeToExecute,
		float64(len(ports))/timeToExecute.Seconds(),
	)

	// show open ports
	for _, scanResult := range scanResults {
		if scanResult.IsOpen {
			fmt.Printf("Open port found: %#v\n", scanResult)
		}
	}
}
