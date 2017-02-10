package csmsg

import (
	"gameproject/common"
	"gameproject/server/db/table"
	"log"
)

type CAddLevelProcess struct {
	msg   *CAddLevel
	trans *common.Trans
}

func (this *CAddLevelProcess) Process() bool {
	defer func() {
		if err := recover(); err != nil {
			log.Println("CAddLevelProcess Error:", err)
		}
	}()

	sendInfo := &SLevelInfo{}
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
