package message

import (
	"gameproject/server/cacheMgr"
	"gameproject/server/protocol"
	"log"
)

type CCreateRoleProcess struct {
	CCreateRole
}

func (this *CCreateRoleProcess) Process() {
	k := "NAME" + this.Name
	v := cacheMgr.GetKV(k)

	sendInfo := &protocol.SCreateRole{}
	if v != "" {
		sendInfo.Res = protocol.SCreateRole_NAME_DUPLICATED
	}
	sendInfo.Res = protocol.SCreateRole_SUCCESS
	sendInfo.Info = &protocol.SRoleList_RoleInfo{}
	sendInfo.Info.RoleId = 123456789
	sendInfo.Info.RoleName = this.Name
	sendInfo.Info.Level = 1
	sendInfo.Info.School = this.School
	sendInfo.Info.ShowFashion = true
	err := this.Send(sendInfo)
	if err != nil {
		log.Panic("CCreateRoleProcess Send SCreateRole error:", err)
	}
}
