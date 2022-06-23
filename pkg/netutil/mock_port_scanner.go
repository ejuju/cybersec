package netutil

type MockPortScanner struct {
	PortBanner string
	PortIsOpen bool
}

func (mockPS *MockPortScanner) Scan(host string, port int, network string) (string, bool) {
	return mockPS.PortBanner, mockPS.PortIsOpen
}
