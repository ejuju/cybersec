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
