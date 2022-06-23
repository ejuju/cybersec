package netutil

// PortScanner can scan a port on a network for a given host.
// It returns a possible banner (string read from port connection)
// and a bool (true if the connection is open)
type PortScanner interface {
	Scan(host string, port int, network string) (string, bool)
}
