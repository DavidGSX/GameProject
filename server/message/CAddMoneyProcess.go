package message

import (
	"gameproject/server/lockMgr"
	"gameproject/server/table"
	"log"
	"strconv"
)

type CAddMoneyProcess struct {
	CAddMoney
}

func (this *CAddMoneyProcess) Process() {
	k := "MONEY_" + strconv.FormatUint(this.RoleId, 10)
	lockMgr.Lock(k)
	defer lockMgr.Unlock(k)

	sendInfo := &SMoneyInfo{}
	sendInfo.RoleId = this.RoleId

	v := table.GetProperty(this.RoleId)
	if v == nil {
		log.Panic("CAddMoneyProcess Role Not Exist RoleId:", this.RoleId)
	} else {
		v.Money += this.Num
	}
	v.Save()

	sendInfo.Total += v.Money
	err := this.Send(sendInfo)
	if err != nil {
		log.Panic("CAddMoneyProcess Send SMoneyInfo error:", err)
	}
}
