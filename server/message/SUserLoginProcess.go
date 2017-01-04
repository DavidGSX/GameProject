package message

import (
	"log"
)

type SUserLoginProcess struct {
	msg *SUserLogin
}

func (this *SUserLoginProcess) Process(msg *SUserLogin) {
	this.msg = msg
	log.Println("to do SUserLoginProcess")
}
