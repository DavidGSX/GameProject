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