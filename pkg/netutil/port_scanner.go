package netutil

import (
	"net"
	"strconv"
	"time"
)

// PortScanner can scan a port on a network for a given host.
// It returns a possible banner (string read from port connection)
// and a bool (true if the connection is open)
type PortScanner interface {
	Scan(host string, port int, network string) (string, bool)
}

// DefaultPortScanner is a simple port scanner implementation
type DefaultPortScanner struct {
	DialTimeout      time.Duration
	ReadTimeout      time.Duration
	BannerBufferSize int
}

// Host is the IP address or domain name.
// Port is the UDP / TCP port.
// Network can be: "tcp", "tcp4" (IPv4-only), "tcp6" (IPv6-only), "udp", "udp4" (IPv4-only), "udp6" (IPv6-only), "ip", "ip4" (IPv4-only), "ip6" (IPv6-only), "unix", "unixgram" and "unixpacket".
func (sps *DefaultPortScanner) Scan(host string, port int, network string) (string, bool) {
	// dial target with given protocol
	// conn, err := net.DialTimeout(proto, host+":"+strconv.Itoa(port), sps.DialTimeout)
	conn, err := net.DialTimeout(network, host+":"+strconv.Itoa(port), sps.DialTimeout)
	if err != nil {
		return "", false
	}
	defer conn.Close()

	// try to read from connection
	buf := make([]byte, sps.BannerBufferSize)
	conn.SetReadDeadline(time.Now().Add(sps.ReadTimeout))
	n, err := conn.Read(buf)
	if err != nil {
		return "", true
	}

	return string(buf[:n]), true
}
