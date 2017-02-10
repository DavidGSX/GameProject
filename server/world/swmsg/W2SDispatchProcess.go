package swmsg

import (
	"gameproject/common"
	"log"
)

type W2SDispatchProcess struct {
	msg   *W2SDispatch
	trans *common.Trans
}

func (this *W2SDispatchProcess) Process() bool {
	defer func() {
		if err := recover(); err != nil {
			log.Println("W2SDispatchProcess Error:", err)
		}
	}()
	
	log.Println("to do W2SDispatchProcess")
	return true
}
