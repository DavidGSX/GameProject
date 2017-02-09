package csmsg

import (
	"gameproject/common"
	"gameproject/server/db/table"
	"log"
)

type CAddMoneyProcess struct {
	msg   *CAddMoney
	trans *common.Trans
}

func (this *CAddMoneyProcess) Process() bool {
	defer func() {
		if err := recover(); err != nil {
			log.Println("CAddMoneyProcess Error:", err)
		}
	}()

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
