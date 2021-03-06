package puck

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

	logEntry.SetLevel("debug")
	// logEntry.SetLevel("fata") // throw a panic

	ctx = logEntry.WrapContextLogger(ctx)

	GetLogger(ctx).Debugf("%s is so cute", "catty")

	GetLogger(ctx).Info("doggy is also cute!")
}

func TestFieldLog(t *testing.T) {
	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()

	logEntry := NewLogger()

	logEntry.SetLevel("info")

	// right now not in sequence, maybe ordered by key later
	logEntry.SetField("center", "yellow")
	logEntry.SetField("tail", "red")
	logEntry.SetField("head", "green")
	logEntry.SetField("traceId", fmt.Sprint(randInt(5000)))

	ctx = logEntry.WrapContextLogger(ctx)

	GetLogger(ctx).Warnf("%s is so sweet", "apple")

	GetLogger(ctx).Error("banana is also sweet!")
}

func randInt(x int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(x)
}
