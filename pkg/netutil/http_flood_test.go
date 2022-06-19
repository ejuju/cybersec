package netutil

import (
	"context"
	"net"
	"net/http"
	"sync"
	"testing"
	"time"
)

func TestHTTPFlood_Attack(t *testing.T) {
	maxDuration := 50 * time.Millisecond
	var numRequests uint64 = 2
	received := 0
	mu := &sync.Mutex{}
	useragent := ""

	router := http.NewServeMux()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		received++
		mu.Unlock()
		useragent = r.UserAgent()
	})

	serverctx, cancelserver := context.WithCancel(context.Background())
	defer cancelserver()

	server := http.Server{
		Handler:     router,
		Addr:        ":8080",
		BaseContext: func(net.Listener) context.Context { return serverctx },
		ConnContext: func(ctx context.Context, c net.Conn) context.Context { return serverctx },
	}

	// listen and serve in seperate routine
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			t.Error(err)
			return
		}
		return
	}()

	// wait for server to start
	time.Sleep(50 * time.Millisecond)

	ctx, cancel := context.WithTimeout(context.Background(), maxDuration)
	defer cancel()

	err := (&HTTPFlood{
		Context:       ctx,
		RequestURL:    "http://localhost:8080/",
		MaxRequests:   numRequests,
		RequestMethod: http.MethodGet,
		RequestBody:   nil,
		NumGoroutines: 1,
	}).Attack()
	if err != nil {
		t.Fatal(err)
	}

	// stop server
	cancelserver()

	t.Run("should stop the attack after deadline is exceeded", func(t *testing.T) {
		if deadline, _ := ctx.Deadline(); deadline.Before(time.Now()) {
			t.Fatal()
		}
	})

	t.Run("should send the right number of requests", func(t *testing.T) {
		if received != int(numRequests) {
			t.Fatalf("unexpected number of requests, want: %v, but got: %v\n", numRequests, received)
		}
	})

	t.Run("should send a request with user-agent", func(t *testing.T) {
		if useragent == "" {
			t.Fatalf("empty user-agent")
		}
	})
}
