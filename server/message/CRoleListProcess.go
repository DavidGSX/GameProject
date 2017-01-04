package message

import (
	"log"
)

type CRoleListProcess struct {
	msg *CRoleList
}

func (this *CRoleListProcess) Process(msg *CRoleList) {
	this.msg = msg
	log.Println("to do CRoleListProcess")
}
