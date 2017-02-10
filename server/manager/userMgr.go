package manager

import (
	"gameproject/common"
	"gameproject/common/cache"
	"gameproject/server/db/table"
	"gameproject/server/world/swmsg"
	"gameproject/world/msgProto"
	"log"
	"time"

	"github.com/golang/protobuf/proto"
)

func OnUserLogout(userId string, roleId uint64) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("OnUserLogout Error:", err)
		}
	}()

	v := table.SelectProperty(roleId)
	if v == nil {
		log.Panic("CAddMoneyProcess Role Not Exist RoleId:", roleId)
	}

	sendInfo := &swmsg.S2WRoleInfoChange{}
	sendInfo.UserId = userId
	sendInfo.Info = &msgProto.S2WRoleInfoChange_RoleInfo{}
	sendInfo.Info.ZoneId = cache.ZoneId
	sendInfo.Info.RoleId = roleId
	sendInfo.Info.RoleName = v.RoleName
	sendInfo.Info.Level = v.Level
	sendInfo.Info.Sex = v.Sex
	sendInfo.Info.School = v.School
	sendInfo.Info.Lasttime = (uint64)(time.Now().Unix())

	data, err := proto.Marshal(sendInfo.GetMsg())
	if err != nil {
		log.Panic("CReqServerRoleInfosProcess msgProto.S2WRoleInfoChange Marshal Error,", err)
	}
	oct := &common.Octets{}
	oct.MarshalUint32(uint32(len(data)))
	oct.MarshalUint32(sendInfo.MsgType())
	oct.MarshalBytesOnly(data)
	GetWorldConn().Send(oct.GetBuf())
}
