package message

import (
	"gameproject/server/msgProto"
	"gameproject/server/table"
	"gameproject/server/transMgr"
	"log"
)

type CRoleListProcess struct {
	msg   *CRoleList
	trans *transMgr.Trans
}

func (this *CRoleListProcess) Process() bool {
	sendInfo := &SRoleList{}
	u := table.SelectUser(this.msg.Getl().GetUserId())
	if u != nil {
		for _, rId := range u.RoleIdList {
			roleInfo := table.SelectProperty(rId)
			if roleInfo != nil {
				r := &msgProto.SRoleList_RoleInfo{}
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
	err := this.msg.Send(sendInfo)
	if err != nil {
		log.Panic("CRoleListProcess Send SRoleList error:", err)
	}
	return true
}
