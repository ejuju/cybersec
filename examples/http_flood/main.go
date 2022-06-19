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
		Ctx:           ctx,
		RequestURL:    "http://localhost:8080",
		RequestMethod: http.MethodGet,
		RequestBody:   nil,
	}

	err := httptarget.Attack()
	if err != nil {
		log.Fatal(err)
	}
}
