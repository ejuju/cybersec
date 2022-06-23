package netutil

import (
	"testing"

	"github.com/ejuju/cybersec/internal/testutil"
)

func TestScanPorts(t *testing.T) {
	t.Run("should return one result per port and per network", func(t *testing.T) {
		numPorts := 10
		ports := PortRange(0, numPorts)
		networks := []string{"tcp", "udp"}

		results := ScanPorts(PortScanConfiguration{
			Ports:    ports,
			Networks: AllNetworks,
			Scanner:  &MockPortScanner{},
		})

		testutil.Check(t, testutil.CheckNotEqualError(len(results), numPorts*len(networks)))
	})
}
