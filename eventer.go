package paw

import (
	"github.com/johnhaha/echo"
	"github.com/johnhaha/nose"
)

func onLog(c *echo.SubCtx) {
	msg := c.Data
	n := nose.NewPageClient(notionToken, logPageID)
	n.AppendTextBlock(getNowTimeString() + " | " + msg)
}
