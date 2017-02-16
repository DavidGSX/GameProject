package csproc

import (
	"gameproject/common"
	"gameproject/server/client/csmsg"
	"gameproject/server/client/msgMgr"
	"log"
)

type CEnterWorldProcess struct {
	msg   *csmsg.CEnterWorld
	trans *common.Trans
}

func (this *CEnterWorldProcess) Clone() msgMgr.IProcess {
	return new(CEnterWorldProcess)
}

func (this *CEnterWorldProcess) SetMsg(m msgMgr.MsgInfo) {
	this.msg = m.(*csmsg.CEnterWorld)
}

func (this *CEnterWorldProcess) SetTrans(t *common.Trans) {
	this.trans = t
}

func (this *CEnterWorldProcess) Process() bool {
	defer func() {
		if err := recover(); err != nil {
			log.Println("CEnterWorldProcess Error:", err)
		}
	}()

	roleId := this.msg.RoleId
	this.msg.Getl().SetRoleId(roleId)

	sendInfo := &csmsg.SEnterWorld{}
	sendInfo.RoleId = roleId
	err := this.msg.Send2Link(sendInfo)
	if err != nil {
		log.Panic("CEnterWorldProcess Send SEnterWorld error:", err)
	}

	return true
}
