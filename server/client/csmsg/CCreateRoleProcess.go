package csmsg

import (
	"gameproject/common"
	"gameproject/server/client/csproto"
	"gameproject/server/db/cacheMgr"
	"gameproject/server/db/table"
	"gameproject/server/rpcMgr"
	"log"
	"time"
)

type CCreateRoleProcess struct {
	msg   *CCreateRole
	trans *common.Trans
}

func (this *CCreateRoleProcess) Process() bool {
	defer func() {
		if err := recover(); err != nil {
			log.Println("CCreateRoleProcess Error:", err)
		}
	}()

	sendInfo := &SCreateRole{}
	var rId uint64
	if rpcMgr.NameExist(this.msg.Name) {
		sendInfo.Res = csproto.SCreateRole_NAME_DUPLICATED
	} else {
		rId = cacheMgr.GetNextRoleId()
		rpcMgr.NameInsert(this.msg.Name)
		sendInfo.Res = csproto.SCreateRole_SUCCESS
	}

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

	sendInfo.Info = &csproto.SRoleList_RoleInfo{}
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
