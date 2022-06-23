package netutil

// AllNetworks is a set of the various networks over which the scan can be performed
var AllNetworks = map[string]struct{}{
	"tcp4":       {},
	"tcp6":       {},
	"udp4":       {},
	"udp6":       {},
	"ip4":        {},
	"ip6":        {},
	"unix":       {},
	"unixgram":   {},
	"unixpacket": {},
}
