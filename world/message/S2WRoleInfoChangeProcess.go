package message

import (
	"gameproject/common"
	"gameproject/world/db/dbProto"
	"gameproject/world/db/table"
	"log"
)

type S2WRoleInfoChangeProcess struct {
	msg   *S2WRoleInfoChange
	trans *common.Trans
}

func (this *S2WRoleInfoChangeProcess) Process() bool {
	defer func() {
		if err := recover(); err != nil {
			log.Println("S2WRoleInfoChangeProcess Error:", err)
		}
	}()

	vall := table.GetAllRoleInfo(this.trans, this.msg.UserId)
	if vall == nil {
		vall = table.NewAllRoleInfo(this.trans, this.msg.UserId)
	}

	vrole, okr := vall.ZoneId2Info[this.msg.Info.ZoneId]
	if okr == false {
		vrole = &dbProto.AllRoleInfo_ServerRoleInfo{}
		vall.ZoneId2Info = make(map[uint32]*dbProto.AllRoleInfo_ServerRoleInfo)
		vall.ZoneId2Info[this.msg.Info.ZoneId] = vrole
	}

	info, oki := vrole.RoleId2Info[this.msg.Info.RoleId]
	if oki == false {
		info = &dbProto.AllRoleInfo_RoleInfo{}
		vrole.RoleId2Info = make(map[uint64]*dbProto.AllRoleInfo_RoleInfo)
		vrole.RoleId2Info[this.msg.Info.RoleId] = info
	}

	info.RoleName = this.msg.Info.RoleName
	info.Level = this.msg.Info.Level
	info.School = this.msg.Info.School
	info.Sex = this.msg.Info.Sex
	info.Lasttime = this.msg.Info.Lasttime

	return true
}
