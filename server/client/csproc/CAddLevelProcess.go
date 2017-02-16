package csproc

import (
	"gameproject/common"
	"gameproject/server/client/csmsg"
	"gameproject/server/client/msgMgr"
	"gameproject/server/db/table"
	"log"
)

type CAddLevelProcess struct {
	msg   *csmsg.CAddLevel
	trans *common.Trans
}

func (this *CAddLevelProcess) Clone() msgMgr.IProcess {
	return new(CAddLevelProcess)
}

func (this *CAddLevelProcess) SetMsg(m msgMgr.MsgInfo) {
	this.msg = m.(*csmsg.CAddLevel)
}

func (this *CAddLevelProcess) SetTrans(t *common.Trans) {
	this.trans = t
}

func (this *CAddLevelProcess) Process() bool {
	defer func() {
		if err := recover(); err != nil {
			log.Println("CAddLevelProcess Error:", err)
		}
	}()

	sendInfo := &csmsg.SLevelInfo{}
	roleId := this.msg.Getr()

	v := table.GetProperty(this.trans, roleId)
	if v == nil {
		log.Panic("CAddLevelProcess Role Not Exist RoleId:", roleId)
	} else {
		v.Level += this.msg.Num
	}

	sendInfo.Level = v.Level
	err := this.msg.Send2Link(sendInfo)
	if err != nil {
		log.Panic("CAddLevelProcess Send SLevelInfo error:", err)
	}
	return true
}
