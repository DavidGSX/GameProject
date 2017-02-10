package swmsg

import (
	"gameproject/common"
	"log"
)

type W2SServerStartResProcess struct {
	msg   *W2SServerStartRes
	trans *common.Trans
}

func (this *W2SServerStartResProcess) Process() bool {
	defer func() {
		if err := recover(); err != nil {
			log.Println("W2SServerStartResProcess Error:", err)
		}
	}()

	log.Println("W2SServerStartResProcess Received!")
	return true
}
