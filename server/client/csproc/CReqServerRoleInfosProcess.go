package csproc

import (
	"gameproject/common"
	"gameproject/common/cache"
	"gameproject/server/client/csmsg"
	"gameproject/server/client/msgMgr"
	"gameproject/server/world/swmsg"
	"gameproject/world/msgProto"
	"log"

	"github.com/golang/protobuf/proto"
)

type CReqServerRoleInfosProcess struct {
	msg   *csmsg.CReqServerRoleInfos
	trans *common.Trans
}

func (this *CReqServerRoleInfosProcess) Clone() msgMgr.IProcess {
	return new(CReqServerRoleInfosProcess)
}

func (this *CReqServerRoleInfosProcess) SetMsg(m msgMgr.MsgInfo) {
	this.msg = m.(*csmsg.CReqServerRoleInfos)
}

func (this *CReqServerRoleInfosProcess) SetTrans(t *common.Trans) {
	this.trans = t
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
