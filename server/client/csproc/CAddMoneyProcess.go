package csproc

import (
	"gameproject/common"
	"gameproject/server/client/csmsg"
	"gameproject/server/client/msgMgr"
	"gameproject/server/db/table"
	"log"
)

type CAddMoneyProcess struct {
	msg   *csmsg.CAddMoney
	trans *common.Trans
}

func (this *CAddMoneyProcess) Clone() msgMgr.IProcess {
	return new(CAddMoneyProcess)
}

func (this *CAddMoneyProcess) SetMsg(m msgMgr.MsgInfo) {
	this.msg = m.(*csmsg.CAddMoney)
}

func (this *CAddMoneyProcess) SetTrans(t *common.Trans) {
	this.trans = t
}

func (this *CAddMoneyProcess) Process() bool {
	defer func() {
		if err := recover(); err != nil {
			log.Println("CAddMoneyProcess Error:", err)
		}
	}()

	sendInfo := &csmsg.SMoneyInfo{}
	roleId := this.msg.Getr()

	v := table.GetProperty(this.trans, roleId)
	if v == nil {
		log.Panic("CAddMoneyProcess Role Not Exist RoleId:", roleId)
	} else {
		v.Money += this.msg.Num
	}

	sendInfo.Money = v.Money
	err := this.msg.Send2Link(sendInfo)
	if err != nil {
		log.Panic("CAddMoneyProcess Send SMoneyInfo error:", err)
	}
	return true
}
