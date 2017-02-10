package csmsg

import (
	"gameproject/common"
	"gameproject/common/cache"
	"gameproject/server/world/swmsg"
	"gameproject/world/msgProto"
	"log"

	"github.com/golang/protobuf/proto"
)

type CReqServerRoleInfosProcess struct {
	msg   *CReqServerRoleInfos
	trans *common.Trans
}

func (this *CReqServerRoleInfosProcess) Process() bool {
	defer func() {
		if err := recover(); err != nil {
			log.Println("CReqServerRoleInfosProcess Error:", err)
		}
	}()

	sendInfo := &swmsg.S2WRoleInfoRequest{}
	sendInfo.ZoneId = cache.ZoneId
	sendInfo.UserId = this.msg.Getl().GetUserId()
	sendInfo.RoleId = this.msg.Getr()
	sendInfo.Req = msgProto.S2WRoleInfoRequest_ROLE_LIST
	data, err := proto.Marshal(sendInfo.GetMsg())
	if err != nil {
		log.Panic("CReqServerRoleInfosProcess msgProto.S2WRoleInfoRequest Marshal Error,", err)
	}
	oct := &common.Octets{}
	oct.MarshalUint32(uint32(len(data)))
	oct.MarshalUint32(sendInfo.MsgType())
	oct.MarshalBytesOnly(data)
	this.msg.Getw().Send(oct.GetBuf())
	return true
}
