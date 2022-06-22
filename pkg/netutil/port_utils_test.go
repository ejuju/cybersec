package netutil

import "testing"

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
