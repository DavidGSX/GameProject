package message

import (
	"gameproject/world/transMgr"
	"log"
)

type GS2WServerStartProcess struct {
	msg   *GS2WServerStart
	trans *transMgr.Trans
}

func (this *GS2WServerStartProcess) Process() bool {
	log.Println("to do GS2WServerStartProcess")
	return true
}
