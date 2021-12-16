package paw

import (
	"context"

	"github.com/johnhaha/echo"
)

func startListen(ctx context.Context) {
	e := echo.NewSuber()
	e.Add(logChannel, onLog)
	e.Sub(ctx)
}
