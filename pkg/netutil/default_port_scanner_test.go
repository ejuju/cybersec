package netutil

import (
	"net"
	"strconv"
	"testing"
	"time"
)

func TestDefaultPortScanner_Scan(t *testing.T) {
	t.Parallel()

	udpTestPort := 58007
	tcpTestPort := 58008
	closedPort := 0
	testBanner := "abc"

	// open tcp port
	go func() {
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
			conn.Write([]byte(testBanner))
			defer conn.Close()
		}
	}()

	// open udp port
	go func() {
		pc, err := net.ListenPacket("udp", ":"+strconv.Itoa(udpTestPort))
		if err != nil {
			panic(err)
		}
		defer pc.Close()

		for {
		}
	}()

	// quick and dirty to wait for listeners to start
	time.Sleep(50 * time.Millisecond)

	scanner := &DefaultPortScanner{
		DialTimeout:      50 * time.Millisecond,
		ReadTimeout:      50 * time.Millisecond,
		BannerBufferSize: len(testBanner),
	}

	t.Run("should detect open TCP ports and read their banner", func(t *testing.T) {
		banner, isopen := scanner.Scan("localhost", tcpTestPort, "tcp")
		if !isopen {
			t.Fatalf("TCP port %d is closed (according to the scan)", tcpTestPort)
		}
		if banner != testBanner {
			t.Fatalf("unexpected banner: want %#v, but got %#v", testBanner, banner)
		}

		_, isopen = scanner.Scan("localhost", closedPort, "tcp")
		if isopen {
			t.Fatalf("TCP port %d is open (according to the scan)", closedPort)
		}
	})

	t.Run("should detect open UDP ports", func(t *testing.T) {
		_, isopen := scanner.Scan("localhost", udpTestPort, "udp")
		if !isopen {
			t.Fatalf("UDP port %d is closed (according to the scan)", udpTestPort)
		}

		_, isopen = scanner.Scan("localhost", closedPort, "udp")
		if isopen {
			t.Fatalf("UDP port %d is open (according to the scan)", closedPort)
		}
	})
}
