package paw

import (
	"context"

	"github.com/johnhaha/echo"
)

var (
	logPubSub = echo.NewPubSub[LogMsg]()
)

func startListen(ctx context.Context) {
	logPubSub.Sub(ctx, appName, 10, onLog)
}
