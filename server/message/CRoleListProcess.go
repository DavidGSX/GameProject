package message

import (
	"gameproject/server/msgProto"
	"gameproject/server/table"
	"log"
)

type CRoleListProcess struct {
	CRoleList
}

func (this *CRoleListProcess) Process() {
	sendInfo := &SRoleList{}
	u := table.GetUser(this.Getl().GetUserId())
	if u != nil {
		for _, rId := range u.RoleIdList {
			roleInfo := table.GetProperty(rId)
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
	err := this.Send(sendInfo)
	if err != nil {
		log.Panic("CRoleListProcess Send SRoleList error:", err)
	}
}
