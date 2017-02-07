package message

import (
	"gameproject/world/transMgr"
	"log"
)

type GS2WSendInfoProcess struct {
	msg   *GS2WSendInfo
	trans *transMgr.Trans
}

func (this *GS2WSendInfoProcess) Process() bool {
	log.Println("to do GS2WSendInfoProcess")
	return true
}
