package paw

import (
	"log"

	"github.com/johnhaha/echo"
	"github.com/johnhaha/nose"
)

func onLog(c *echo.SubCtx) {
	msg := c.Data
	n := nose.NewPageClient(notionToken, logPageID)
	err := n.AppendTextBlock(getNowTimeString() + " | " + msg)
	if err != nil {
		log.Println("paw log failed", err.Error(), msg)
	}
}
