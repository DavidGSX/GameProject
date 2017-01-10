package message

import (
	"gameproject/server/cacheMgr"
	"gameproject/server/msgProto"
	"gameproject/server/rpcMgr"
	"gameproject/server/table"
	"log"
	"time"
)

type CCreateRoleProcess struct {
	CCreateRole
}

func (this *CCreateRoleProcess) Process() {
	sendInfo := &SCreateRole{}
	var rId uint64
	if rpcMgr.NameExist(this.Name) {
		sendInfo.Res = msgProto.SCreateRole_NAME_DUPLICATED
	} else {
		rId = cacheMgr.GetNextRoleId()
		rpcMgr.NameInsert(this.Name)
	}

	/*
		t := table.GetName(this.Name)
		if t != nil {
			sendInfo.Res = msgProto.SCreateRole_NAME_DUPLICATED
		} else {
			t = table.NewName(this.Name)
			t.UserId = this.Getl().GetUserId()
			rId = cacheMgr.GetNextRoleId()
			t.RoleId = rId
			t.CreateTime = uint64(time.Now().Unix())
			t.Save()
		}
	*/

	p := table.GetProperty(rId)
	if p != nil {
		log.Panic("CCreateRoleProcess Role Id Duplicate rId:", rId)
	}
	p = table.NewProperty(rId)
	p.UserId = this.Getl().GetUserId()
	p.RoleName = this.Name
	p.CreateTime = uint64(time.Now().Unix())
	p.School = this.School
	p.Sex = this.Sex
	p.Level = 1
	p.Save()

	u := table.GetUser(this.Getl().GetUserId())
	if u == nil {
		u = table.NewUser(this.Getl().GetUserId())
	}
	u.RoleIdList = append(u.RoleIdList, rId)
	u.LastLoginRoleId = rId
	u.CreateTime = uint64(time.Now().Unix())
	u.Save()

	sendInfo.Res = msgProto.SCreateRole_SUCCESS
	sendInfo.Info = &msgProto.SRoleList_RoleInfo{}
	sendInfo.Info.RoleId = rId
	sendInfo.Info.RoleName = this.Name
	sendInfo.Info.Level = 1
	sendInfo.Info.School = this.School
	sendInfo.Info.ShowFashion = true
	err := this.Send(sendInfo)
	if err != nil {
		log.Panic("CCreateRoleProcess Send SCreateRole error:", err)
	}
}
