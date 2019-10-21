package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/minipub/puck"
)

var (
	xtime = 5
	xcnt  = 3
)

func main() {
	addr := "127.0.0.1:9090"

	http.HandleFunc("/test", testHandler)
	log.Printf("Server start: %s", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()

	logEntry := puck.NewLogger()

	logEntry.SetLevel("info")

	logEntry.Set("traceId", fmt.Sprint(randInt(5000)))

	ctx = logEntry.WrapContextLogger(ctx)

	xcnt = xcnt - 1
	xsleep := xcnt * xtime

	puck.GetLogger(ctx).Infof("before: xcnt[ %d ], xsleep[ %d ]", xcnt, xsleep)
	time.Sleep(time.Duration(xsleep) * time.Second)
	puck.GetLogger(ctx).Infof("after: xcnt[ %d ], xsleep[ %d ]", xcnt, xsleep)
}

func randInt(x int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(x)
}
