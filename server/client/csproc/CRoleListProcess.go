package csproc

import (
	"gameproject/common"
	"gameproject/server/client/csmsg"
	"gameproject/server/client/csproto"
	"gameproject/server/client/msgMgr"
	"gameproject/server/db/table"
	"log"
)

type CRoleListProcess struct {
	msg   *csmsg.CRoleList
	trans *common.Trans
}

func (this *CRoleListProcess) Clone() msgMgr.IProcess {
	return new(CRoleListProcess)
}

func (this *CRoleListProcess) SetMsg(m msgMgr.MsgInfo) {
	this.msg = m.(*csmsg.CRoleList)
}

func (this *CRoleListProcess) SetTrans(t *common.Trans) {
	this.trans = t
}

func (this *CRoleListProcess) Process() bool {
	defer func() {
		if err := recover(); err != nil {
			log.Println("CRoleListProcess Error:", err)
		}
	}()

	sendInfo := &csmsg.SRoleList{}
	u := table.SelectUser(this.msg.Getl().GetUserId())
	if u != nil {
		for _, rId := range u.RoleIdList {
			roleInfo := table.SelectProperty(rId)
			if roleInfo != nil {
				r := &csproto.SRoleList_RoleInfo{}
				r.RoleId = rId
				r.RoleName = roleInfo.RoleName
				r.Level = roleInfo.Level
				r.School = roleInfo.School
				sendInfo.Roles = append(sendInfo.Roles, r)
			}
		}
		sendInfo.PreLoginRoleId = u.LastLoginRoleId
	} else {
		sendInfo.PreLoginRoleId = 1
	}
	err := this.msg.Send2Link(sendInfo)
	if err != nil {
		log.Panic("CRoleListProcess Send SRoleList error:", err)
	}
	return true
}
