package swmsg

import (
	"gameproject/common"
	"gameproject/server/client/csmsg"
	"gameproject/server/client/csproto"
	"gameproject/server/client/msgMgr"
	"log"
)

type W2SRoleInfoResponseProcess struct {
	msg   *W2SRoleInfoResponse
	trans *common.Trans
}

func (this *W2SRoleInfoResponseProcess) Process() bool {
	defer func() {
		if err := recover(); err != nil {
			log.Println("W2SRoleInfoResponseProcess Error:", err)
		}
	}()

	sendInfo := &csmsg.SServerRoleInfos{}
	sendInfo.RoleId = this.msg.RoleId

	if this.msg.Info != nil {
		sendInfo.Info = make([]*csproto.SServerRoleInfos_RoleInfo, 0)
		for _, rInfo := range this.msg.Info {
			info := &csproto.SServerRoleInfos_RoleInfo{}
			info.ZoneId = rInfo.ZoneId
			info.RoleId = rInfo.RoleId
			info.RoleName = rInfo.RoleName
			info.Level = rInfo.Level
			info.School = rInfo.School
			info.Sex = rInfo.Sex
			info.Lasttime = rInfo.Lasttime
			sendInfo.Info = append(sendInfo.Info, info)
		}
	}

	data, err := msgMgr.MarshalMsg(sendInfo)
	if err != nil {
		log.Panic("Marshal SServerRoleInfos Error,", err)
	}
	this.msg.Gets().SendByRoleIds([]uint64{sendInfo.RoleId}, data)
	return true
}
