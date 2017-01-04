package message

import (
	"log"
)

type SRoleListProcess struct {
	msg *SRoleList
}

func (this *SRoleListProcess) Process(msg *SRoleList) {
	this.msg = msg
	log.Println("to do SRoleListProcess")
}
