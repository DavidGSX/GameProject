package message

import (
	"gameproject/server/table"
	"gameproject/server/transMgr"
	"log"
)

type CAddMoneyProcess struct {
	msg   *CAddMoney
	trans *transMgr.Trans
}

func (this *CAddMoneyProcess) Process() bool {
	sendInfo := &SMoneyInfo{}
	sendInfo.RoleId = this.msg.RoleId

	v := table.GetProperty(this.trans, this.msg.RoleId)
	if v == nil {
		log.Panic("CAddMoneyProcess Role Not Exist RoleId:", this.msg.RoleId)
	} else {
		v.Money += this.msg.Num
	}

	sendInfo.Total = v.Money
	err := this.msg.Send(sendInfo)
	if err != nil {
		log.Panic("CAddMoneyProcess Send SMoneyInfo error:", err)
	}
	return true
}
