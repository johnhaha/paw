package paw

import (
	"log"

	"github.com/johnhaha/echo"
)

//listen to log
func onLog(c *echo.SubCtx) {
	var msg LogMsg
	err := c.Parser(&msg)
	if err != nil {
		log.Println("paw parse log msg failed")
		return
	}
	err = msg.Log()
	if err != nil {
		log.Println("paw log failed", err.Error(), msg)
	}
}
