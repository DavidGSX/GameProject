package message

import (
	"gameproject/common"
	"gameproject/server/cacheMgr"
	"gameproject/server/protocol"
	"log"

	"github.com/golang/protobuf/proto"
)

type CCreateRoleProcess struct {
	msg *CCreateRole
}

func (this *CCreateRoleProcess) Process(msg *CCreateRole) {
	this.msg = msg

	k := "NAME" + msg.Proto.Name
	v := cacheMgr.GetKV(k)

	sendInfo := &protocol.SCreateRole{}
	if v != "" {
		sendInfo.Res = protocol.SCreateRole_NAME_DUPLICATED
	}
	sendInfo.Res = protocol.SCreateRole_SUCCESS
	sendInfo.Info = &protocol.SRoleList_RoleInfo{}
	sendInfo.Info.RoleId = 123456789
	sendInfo.Info.RoleName = msg.Proto.Name
	sendInfo.Info.Level = 1
	sendInfo.Info.School = msg.Proto.School
	sendInfo.Info.ShowFashion = true
	data, err := proto.Marshal(sendInfo)
	if err != nil {
		log.Panic("marshal error:", err)
	}
	oct := &common.Octets{}
	oct.MarshalUint32(uint32(len(data)))
	oct.MarshalUint32(1006)
	oct.MarshalBytesOnly(data)
	msg.Link.Send(oct.GetBuf())
}
