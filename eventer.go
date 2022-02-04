package paw

import (
	"log"
)

//listen to log
func onLog(msg LogMsg) {
	err := msg.Log()
	if err != nil {
		log.Println("paw log failed", err.Error(), msg)
	}
}
