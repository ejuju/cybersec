package netutil

import (
	"io"
	"net"
	"time"

	"golang.org/x/crypto/ssh"
)

type SSHConnection struct {
	config SSHConnectionConfig
	client *ssh.Client
}

type SSHConnectionConfig struct {
	Address     *net.TCPAddr
	User        string
	Password    string
	DialTimeout time.Duration
}

// Returns an error if the credentials are not right
// Don't forget to call SSHConnection.Close()
func NewSSHConnection(config SSHConnectionConfig) (*SSHConnection, error) {
	// now attempt to login with ssh, if an error occurs,
	// (it shouldn't be a network error since this was checked above)
	clientconfig := &ssh.ClientConfig{
		User:            config.User,
		Auth:            []ssh.AuthMethod{ssh.Password(config.Password)},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         config.DialTimeout,
	}
	clientconfig.SetDefaults()

	// attempt to connect via SSH
	c, err := ssh.Dial("tcp", config.Address.String(), clientconfig)
	// if there is no error (from the call to ssh.Dial), then the credentials are valid
	return &SSHConnection{
		config: config,
		client: c,
	}, err
}

func (conn *SSHConnection) Close() error {
	return conn.client.Close()
}

// returns stdout, stderr and a connection error
func (conn *SSHConnection) Run(command string, stdout, stderr io.Writer) error {
	session, err := conn.client.NewSession()
	if err != nil {
		return err
	}
	defer session.Close()

	session.Stdout = stdout
	session.Stderr = stderr

	err = session.Run(command)
	if err != nil {
		return err
	}

	return err
}
