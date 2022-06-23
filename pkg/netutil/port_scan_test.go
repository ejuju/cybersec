package netutil

import (
	"testing"

	"github.com/ejuju/cybersec/internal/testutil"
)

func TestScanPorts(t *testing.T) {
	t.Run("should return one result per port and per network", func(t *testing.T) {
		numPorts := 10
		ports := PortRange(0, numPorts)
		networks := map[string]struct{}{"tcp": {}, "udp": {}}

		results := ScanPorts(PortScanConfiguration{
			Ports:    ports,
			Networks: networks,
			Scanner:  &MockPortScanner{},
		})

		testutil.Check(t, testutil.CheckNotEqualError(len(results), numPorts*len(networks)))
	})
}

func TestPortRange(t *testing.T) {
	t.Parallel()

	min := 0
	max := 2
	ports := PortRange(min, max)

	t.Run("should a range betwen min (inclusive) and max (exclusive)", func(t *testing.T) {
		// should include minimum
		_, ok := ports[min]
		testutil.Check(t, testutil.CheckNotEqualError(ok, true))

		// should include maximum - 1
		_, ok = ports[max-1]
		testutil.Check(t, testutil.CheckNotEqualError(ok, true))

		// should not include maximum
		_, ok = ports[max]
		testutil.Check(t, testutil.CheckNotEqualError(ok, false))
	})
}
