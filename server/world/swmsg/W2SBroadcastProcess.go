package swmsg

import (
	"gameproject/common"
	"log"
)

type W2SBroadcastProcess struct {
	msg   *W2SBroadcast
	trans *common.Trans
}

func (this *W2SBroadcastProcess) Process() bool {
	defer func() {
		if err := recover(); err != nil {
			log.Println("W2SBroadcastProcess Error:", err)
		}
	}()
	
	log.Println("to do W2SBroadcastProcess")
	return true
}
