package netutil

// PortScanConfiguration holds configuration variables for the ScanPorts function
type PortScanConfiguration struct {
	Host     string
	Ports    map[int]struct{}
	Networks map[string]struct{}
	Scanner  PortScanner
}

// PortScanResult stores information about the outcome of a port scan
// (on a single port and a single network)
type PortScanResult struct {
	Port    int    // port (should be between 0 and 65535)
	Info    string // if port is a common port, the possible purpose is listed here
	Network string // attempted protocol for this port
	IsOpen  bool   // indicates if the port accepts connections
	Banner  string // the banner can provide information about the port
}

// ScanPorts is a utility function to perform port scanning
// over a range of ports and on one or more networks
func ScanPorts(config PortScanConfiguration) []*PortScanResult {
	totalOps := len(config.Ports) * len(config.Networks)
	results := []*PortScanResult{}
	resultsChan := make(chan *PortScanResult, totalOps)
	defer close(resultsChan)

	for port := range config.Ports {
		for network := range config.Networks {
			go func(resultsChan chan<- *PortScanResult, port int, network string) {
				// scan port and append result
				banner, open := config.Scanner.Scan(config.Host, port, network)
				purpose, _ := CommonPorts[port]
				resultsChan <- &PortScanResult{
					Port:    port,
					Network: network,
					IsOpen:  open,
					Banner:  banner,
					Info:    purpose,
				}
			}(resultsChan, port, network)
		}
	}

	for i := 0; i < totalOps; i++ {
		result := <-resultsChan
		results = append(results, result)
	}

	return results
}

// PortRange is a utility function to initialize a continuous set of ports
func PortRange(min, max int) map[int]struct{} {
	ports := map[int]struct{}{}
	for i := min; i < max; i++ {
		ports[i] = struct{}{}
	}
	return ports
}

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

// CommonPorts stores commonly used ports and their associated purpose
var CommonPorts = map[int]string{
	1:    "tcpmux",
	5:    "rje",
	7:    "echo",
	9:    "discard",
	11:   "systat",
	13:   "daytime",
	20:   "ftp-data",
	21:   "ftp",
	22:   "ssh",
	23:   "telnet",
	25:   "smtp",
	43:   "whois",
	53:   "dns",
	67:   "dhcp",
	68:   "dhcp",
	80:   "http",
	110:  "pop3",
	123:  "ntp",
	137:  "netbios",
	138:  "netbios",
	139:  "netbios",
	143:  "imap4",
	443:  "https",
	513:  "rlogin",
	540:  "uucp",
	554:  "rtsp",
	587:  "smtp",
	873:  "rsync",
	902:  "vmware",
	989:  "ftps",
	990:  "ftps",
	1194: "openvpn",
	3306: "mysql",
	5000: "unpn",
	8080: "https-proxy",
	8443: "https-alt",
}
