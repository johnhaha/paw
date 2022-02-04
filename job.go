package paw

import (
	"context"

	"github.com/johnhaha/echo"
)

var (
	logPubSub = echo.NewPubSub[LogMsg]()
)

func startListen(ctx context.Context) {
	logPubSub.Sub(ctx, "paw", 10, onLog)
}
