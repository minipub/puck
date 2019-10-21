package log

import (
	"context"
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestNoFieldLog(t *testing.T) {
	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()

	logEntry := NewLogger()

	logEntry.SetLevel("info")
	// logEntry.SetLevel("fata") // throw a panic

	ctx = logEntry.WrapContextLogger(ctx)

	GetLogger(ctx).Infof("%s is so cute", "catty")

	GetLogger(ctx).Info("doggy is also cute!")
}

func TestFieldLog(t *testing.T) {
	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()

	logEntry := NewLogger()

	logEntry.SetLevel("info")

	// sequence: reversed by key
	logEntry.SetField("center", "yellow")
	logEntry.SetField("tail", "red")
	logEntry.SetField("head", "green")
	logEntry.SetField("traceId", fmt.Sprint(randInt(5000)))

	ctx = logEntry.WrapContextLogger(ctx)

	GetLogger(ctx).Infof("%s is so sweet", "apple")

	GetLogger(ctx).Info("banana is also sweet!")
}

func randInt(x int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(x)
}
