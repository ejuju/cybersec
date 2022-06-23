package netutil

// PortRange is a utility function to initialize a continuous set of ports
func PortRange(min, max int) map[int]struct{} {
	ports := map[int]struct{}{}
	for i := min; i < max; i++ {
		ports[i] = struct{}{}
	}
	return ports
}
