package paw

import (
	"context"

	"github.com/johnhaha/echo"
	"github.com/johnhaha/nose"
)

func StartPaw(ctx context.Context, token string, pageID string, version string) {
	tail := "-log-" + version
	c := nose.NewPageClient(token, pageID)
	res := c.NewEmptyPage(getNowTimeString() + tail)
	logPageID = res
	startListen(ctx)
}

func Log(msg string) error {
	return echo.Pub(logChannel, msg)
}
