package message

import (
	"gameproject/server/cacheMgr"
	"gameproject/server/msgProto"
	"gameproject/server/rpcMgr"
	"gameproject/server/table"
	"gameproject/server/transMgr"
	"log"
	"time"
)

type CCreateRoleProcess struct {
	msg   *CCreateRole
	trans *transMgr.Trans
}

func (this *CCreateRoleProcess) Process() bool {
	sendInfo := &SCreateRole{}
	var rId uint64
	if rpcMgr.NameExist(this.msg.Name) {
		sendInfo.Res = msgProto.SCreateRole_NAME_DUPLICATED
	} else {
		rId = cacheMgr.GetNextRoleId()
		rpcMgr.NameInsert(this.msg.Name)
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

	p := table.GetProperty(this.trans, rId)
	if p != nil {
		log.Panic("CCreateRoleProcess Role Id Duplicate rId:", rId)
	}
	p = table.NewProperty(this.trans, rId)
	p.UserId = this.msg.Getl().GetUserId()
	p.RoleName = this.msg.Name
	p.CreateTime = uint64(time.Now().Unix())
	p.School = this.msg.School
	p.Sex = this.msg.Sex
	p.Level = 1

	u := table.GetUser(this.trans, this.msg.Getl().GetUserId())
	if u == nil {
		u = table.NewUser(this.trans, this.msg.Getl().GetUserId())
	}
	u.RoleIdList = append(u.RoleIdList, rId)
	u.LastLoginRoleId = rId
	u.CreateTime = uint64(time.Now().Unix())

	sendInfo.Res = msgProto.SCreateRole_SUCCESS
	sendInfo.Info = &msgProto.SRoleList_RoleInfo{}
	sendInfo.Info.RoleId = rId
	sendInfo.Info.RoleName = this.msg.Name
	sendInfo.Info.Level = 1
	sendInfo.Info.School = this.msg.School
	sendInfo.Info.ShowFashion = true
	err := this.msg.Send(sendInfo)
	if err != nil {
		log.Panic("CCreateRoleProcess Send SCreateRole error:", err)
	}
	return true
}
