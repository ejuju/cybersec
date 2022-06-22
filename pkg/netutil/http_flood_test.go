package netutil

import (
	"context"
	"net/http"
	"strconv"
	"sync"
	"testing"
	"time"
)

func TestHTTPFlood_Attack(t *testing.T) {
	t.Parallel()

	httpPort := 58001 // make sure this does not conflict with other tests
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

	server := http.Server{
		Handler: router,
		Addr:    ":" + strconv.Itoa((httpPort)),
	}

	// listen and serve in seperate routine
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			t.Error(err)
		}
	}()

	// quick and dirty: wait for server to start
	time.Sleep(50 * time.Millisecond)

	ctx, cancel := context.WithTimeout(context.Background(), maxDuration)
	defer cancel()

	err := (&HTTPFlood{
		Context:       ctx,
		RequestURL:    "http://localhost:" + strconv.Itoa(httpPort) + "/",
		MaxRequests:   numRequests,
		RequestMethod: http.MethodGet,
		RequestBody:   nil,
		NumGoroutines: 2,
	}).Attack()
	if err != nil {
		t.Fatal(err)
	}

	t.Run("should stop the attack after deadline is exceeded", func(t *testing.T) {
		if deadline, _ := ctx.Deadline(); deadline.Before(time.Now()) {
			t.Fatalf("deadline is exceeded by %s", time.Now().Sub(deadline))
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
