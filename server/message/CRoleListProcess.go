package message

import (
	"gameproject/common"
	"gameproject/server/cache"
	"gameproject/server/protocol"
	"log"

	"github.com/golang/protobuf/proto"
)

type CRoleListProcess struct {
	msg *CRoleList
}

func (this *CRoleListProcess) Process(msg *CRoleList) {
	this.msg = msg

	k := "USER" + msg.Link.GetUserId()
	v := cache.GetKV(k)

	sendInfo := &protocol.SRoleList{}
	if v != "" {
		// Decode DB Data
	}
	sendInfo.PreLoginRoleId = 1
	data, err := proto.Marshal(sendInfo)
	if err != nil {
		log.Panic("marshal error:", err)
	}
	oct := &common.Octets{}
	oct.MarshalUint32(uint32(len(data)))
	oct.MarshalUint32(1004)
	oct.MarshalBytesOnly(data)
	msg.Link.Send(oct.GetBuf())
}
