package netutil

import (
	"testing"

	"github.com/ejuju/cybersec/internal/testutil"
)

func TestPortRange(t *testing.T) {
	t.Parallel()

	ports := PortRange(0, 1)

	_, ok := ports[0]
	if !ok {
		t.Fatal("ports should contain 0")
	}

	_, ok = ports[1]
	if ok {
		t.Fatal("ports should not contain 1")
	}
}

func TestCommonPorts(t *testing.T) {
	t.Parallel()
	t.Run("should not be empty", func(t *testing.T) {
		testutil.Check(t, testutil.CheckZeroLengthError[string, int](CommonPorts))
	})
}

func TestAllNetworks(t *testing.T) {
	t.Parallel()
	t.Run("should not be empty", func(t *testing.T) {
		testutil.Check(t, testutil.CheckZeroLengthError[struct{}, string](AllNetworks))
	})
}
