package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/minipub/puck"
	uuid "github.com/satori/go.uuid"
)

func main() {
	addr := "127.0.0.1:9090"

	http.HandleFunc("/test", testHandler)
	log.Printf("Server start: %s", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	var (
		xtime = 5
		xcnt  = 3
	)

	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()

	logEntry := puck.NewLogger()

	logEntry.SetLevel("info")

	logEntry.Set("traceId", genUUIDv4())

	ctx = logEntry.WrapContextLogger(ctx)

	xcnt = xcnt - 1
	xsleep := xcnt * xtime

	GetLogger(ctx).Infof("before: xcnt[ %d ], xsleep[ %d ]", xcnt, xsleep)
	time.Sleep(time.Duration(xsleep) * time.Second)
	GetLogger(ctx).Infof("after: xcnt[ %d ], xsleep[ %d ]", xcnt, xsleep)
}

func genUUIDv4() string {
	return fmt.Sprint(uuid.NewV4())
}
