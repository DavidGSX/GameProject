package message

import (
	"gameproject/server/protocol"

	"github.com/golang/protobuf/proto"
)

type CRoleList struct {
	RoleId uint64
	Proto  protocol.CRoleList
}

func (this *CRoleList) Clone() MsgInfo {
	return new(CRoleList)
}

func (this *CRoleList) SetRoleId(r uint64) {
	this.RoleId = r
}

func (this *CRoleList) Unmarshal(data []byte) error {
	err := proto.Unmarshal(data, &this.Proto)
	return err
}

func (this *CRoleList) Process() {
}

