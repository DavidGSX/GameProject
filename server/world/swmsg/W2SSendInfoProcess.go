package swmsg

import (
	"gameproject/common"
	"log"
)

type W2SSendInfoProcess struct {
	msg   *W2SSendInfo
	trans *common.Trans
}

func (this *W2SSendInfoProcess) Process() bool {
	defer func() {
		if err := recover(); err != nil {
			log.Println("W2SSendInfoProcess Error:", err)
		}
	}()
	
	log.Println("to do W2SSendInfoProcess")
	return true
}
