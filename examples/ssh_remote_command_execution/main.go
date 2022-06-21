package main

import (
	"bytes"
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
		os.Exit(1)
	}
	tcpClient.Close()

	// run SSH command
	user := "admin"
	password := "admin123"

	sshconn, err := netutil.NewSSHConnection(netutil.SSHConnectionConfig{
		Address:     address,
		User:        user,
		Password:    password,
		DialTimeout: 100 * time.Millisecond,
	})

	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}

	command := "uname -a"
	err = sshconn.Run(command, stdout, stderr)
	if err != nil {
		fmt.Printf("unable to run SSH command:\n%#v\nerror: %v\n", command, err)
		os.Exit(1)
	}

	fmt.Printf("SSH command stdout: \"%#v\"\n", stdout)
	fmt.Printf("SSH command stderr: \"%#v\"\n", stderr)
}
