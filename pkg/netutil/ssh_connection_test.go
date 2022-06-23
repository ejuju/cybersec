package netutil

import (
	"bytes"
	"errors"
	"log"
	"net"
	"testing"
	"time"

	"golang.org/x/crypto/ssh"
)

// TestSSHConnection is done by starting a SSH server
func TestSSHConnection(t *testing.T) {
	t.Parallel()
	t.Skip("to fix: unable to connect to ssh server for now")

	serverPort := 2022
	username := "test"
	password := "test123"
	echoString := "hello world"

	// open TCP port for SSH connections
	l, err := net.ListenTCP("tcp", &net.TCPAddr{Port: serverPort})
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		l.Close()
	})

	// accept incoming connections to SSH server
	go startTestSSHServer(t, l, &ssh.ServerConfig{
		PasswordCallback: func(c ssh.ConnMetadata, pwd []byte) (*ssh.Permissions, error) {
			t.Log("hehe", c.User(), pwd, c.SessionID())
			if c.User() != username || string(pwd) != password {
				return nil, errors.New("password rejected for: " + c.User())
			}
			return nil, nil // successful login
		},
	}, echoString)

	time.Sleep(50 * time.Millisecond)

	t.Run("should be able to connect to SSH server", func(t *testing.T) {
		sshconn, err := NewSSHConnection(SSHConnectionConfig{
			Address:  &net.TCPAddr{Port: serverPort},
			User:     username,
			Password: password,
		})
		if err != nil {
			t.Fatal("connection failed", err)
		}
		defer sshconn.Close()

		// check with invalid credentials
		invalidsshconn, err := NewSSHConnection(SSHConnectionConfig{
			Address:  &net.TCPAddr{Port: serverPort},
			User:     "invalid_username",
			Password: "invalid_password",
		})
		if err == nil {
			t.Fatal("connection succeeded but should have failed", err)
		}
		defer invalidsshconn.Close()
	})

	t.Run("should be able to run a command remotely", func(t *testing.T) {
		sshconn, err := NewSSHConnection(SSHConnectionConfig{
			Address:  &net.TCPAddr{Port: serverPort},
			User:     username,
			Password: password,
		})
		if err != nil {
			t.Fatal("connection failed", err)
		}
		defer sshconn.Close()

		buf := &bytes.Buffer{}
		err = sshconn.Run("printf \""+echoString+"\"", buf, nil)
		if err != nil {
			t.Fatal("command execution failed")
		}
		if buf.String() != echoString {
			t.Fatalf("unexpected stdout return string, want %#v, but got %#v", echoString, buf.String())
		}
	})
}

func startTestSSHServer(t *testing.T, l *net.TCPListener, serverConfig *ssh.ServerConfig, echoString string) {
	for {
		// accept incoming TCP connections
		conn, err := l.Accept()
		if err != nil {
			t.Log("failed to accept incoming connection: ", err)
			return
		}
		defer conn.Close()

		// serve SSH
		sshConn, newChannelChan, reqs, err := ssh.NewServerConn(conn, serverConfig)
		if err != nil {
			continue // don't do anything on conn error, this might be normal for the test
		}
		defer sshConn.Close()

		// service shell requests
		// go ssh.DiscardRequests(reqs)
		go func(in <-chan *ssh.Request) {
			for req := range in {
				req.Reply(req.Type == "shell", nil)
			}
		}(reqs)

		// service the incoming ssh channel
		for newChannel := range newChannelChan {
			// check if client channel type is "session"
			if newChannel.ChannelType() != "session" {
				newChannel.Reject(ssh.UnknownChannelType, "unknown channel type")
				continue
			}

			// accept the SSH channel creation request
			sessionChannel, requests, err := newChannel.Accept()
			if err != nil {
				log.Fatalf("Could not accept channel: %v", err)
			}
			defer sessionChannel.Close()

			// handle shell requests
			go func(in <-chan *ssh.Request) {
				for req := range in {
					req.Reply(req.Type == "shell", []byte(echoString))
				}
			}(requests)
		}
	}
}
