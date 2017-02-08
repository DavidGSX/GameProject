package message

import (
	"gameproject/world/transMgr"
	"log"
)

type S2WRoleInfoRequestProcess struct {
	msg   *S2WRoleInfoRequest
	trans *transMgr.Trans
}

func (this *S2WRoleInfoRequestProcess) Process() bool {
	defer func() {
		if err := recover(); err != nil {
			log.Println("S2WRoleInfoRequestProcess Error:", err)
		}
	}()
	
	log.Println("to do S2WRoleInfoRequestProcess")
	return true
}
