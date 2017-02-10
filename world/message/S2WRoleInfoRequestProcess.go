package message

import (
	"gameproject/common"
	"gameproject/world/db/table"
	"gameproject/world/msgProto"
	"log"
)

type S2WRoleInfoRequestProcess struct {
	msg   *S2WRoleInfoRequest
	trans *common.Trans
}

func (this *S2WRoleInfoRequestProcess) Process() bool {
	defer func() {
		if err := recover(); err != nil {
			log.Println("S2WRoleInfoRequestProcess Error:", err)
		}
	}()

	sendInfo := &W2SRoleInfoResponse{}
	sendInfo.RoleId = this.msg.RoleId
	sendInfo.Req = this.msg.Req

	vall := table.GetAllRoleInfo(this.trans, this.msg.UserId)
	if vall.ZoneId2Info != nil {
		sendInfo.Info = make([]*msgProto.S2WRoleInfoChange_RoleInfo, 0)
		for zoneId, zInfo := range vall.ZoneId2Info {
			for roleId, rInfo := range zInfo.RoleId2Info {
				info := &msgProto.S2WRoleInfoChange_RoleInfo{}
				info.ZoneId = zoneId
				info.RoleId = roleId
				info.RoleName = rInfo.RoleName
				info.Level = rInfo.Level
				info.School = rInfo.School
				info.Sex = rInfo.Sex
				info.Lasttime = rInfo.Lasttime
				sendInfo.Info = append(sendInfo.Info, info)
			}
		}
	}

	this.msg.Send(sendInfo)

	return true
}
