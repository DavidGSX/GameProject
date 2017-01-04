package message

import (
	"log"
)

type CUserLoginProcess struct {
	msg *CUserLogin
}

func (this *CUserLoginProcess) Process(msg *CUserLogin) {
	this.msg = msg
	log.Println("to do CUserLoginProcess")
}
