package csmsg

import (
	"gameproject/common"
	"log"
)

type CEnterWorldProcess struct {
	msg   *CEnterWorld
	trans *common.Trans
}

func (this *CEnterWorldProcess) Process() bool {
	defer func() {
		if err := recover(); err != nil {
			log.Println("CEnterWorldProcess Error:", err)
		}
	}()

	roleId := this.msg.RoleId
	this.msg.Getl().SetRoleId(roleId)

	sendInfo := &SEnterWorld{}
	sendInfo.RoleId = roleId
	err := this.msg.Send2Link(sendInfo)
	if err != nil {
		log.Panic("CEnterWorldProcess Send SEnterWorld error:", err)
	}

	return true
}
