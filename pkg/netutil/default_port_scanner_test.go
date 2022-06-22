package netutil

import (
	"net"
	"testing"
	"time"
)

func TestDefaultPortScanner_Scan(t *testing.T) {
	t.Parallel()

	tcpTestPort := 58008
	tcpTestBanner := "abc"

	go func() {
		// open tcp port
		l, err := net.ListenTCP("tcp", &net.TCPAddr{Port: tcpTestPort})
		if err != nil {
			panic(err)
		}
		defer l.Close()

		for {
			conn, err := l.Accept()
			if err != nil {
				panic(err)
			}
			conn.Write([]byte(tcpTestBanner))
			defer conn.Close()
		}
	}()

	scanner := &DefaultPortScanner{
		DialTimeout:      50 * time.Millisecond,
		ReadTimeout:      50 * time.Millisecond,
		BannerBufferSize: len(tcpTestBanner),
	}

	t.Run("should detect open TCP ports", func(t *testing.T) {
		_, isopen := scanner.Scan("localhost", tcpTestPort, "tcp")
		if !isopen {
			t.Fatalf("port %d is closed (according to the scan)", tcpTestPort)
		}
	})

	t.Run("should read the banner of open TCP", func(t *testing.T) {
		banner, _ := scanner.Scan("localhost", tcpTestPort, "tcp")
		if banner != tcpTestBanner {
			t.Fatalf("banner unexpected banner received: want %#v, but got %#v", tcpTestBanner, banner)
		}
	})

	t.Run("should detect closed TCP ports", func(t *testing.T) {
		closedPort := tcpTestPort + 1
		_, isopen := scanner.Scan("localhost", closedPort, "tcp")
		if isopen {
			t.Fatalf("port %d is open (according to the scan)", closedPort)
		}
	})
}
