package paw

import (
	"context"

	"github.com/johnhaha/echo"
	"github.com/johnhaha/nose"
)

func StartPaw(ctx context.Context, token string, pageID string, version string) error {
	tail := "-log-" + version
	c := nose.NewPageClient(token, pageID)
	res, err := c.NewEmptyPage(getNowTimeString() + tail)
	if err != nil {
		return err
	}
	logPageID = res
	notionToken = token
	startListen(ctx)
	return nil
}

func Log(msg string) error {
	return echo.Pub(logChannel, msg)
}
