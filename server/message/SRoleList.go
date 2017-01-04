package message

import (
	"gameproject/server/protocol"

	"github.com/golang/protobuf/proto"
)

type SRoleList struct {
	RoleId uint64
	Proto  protocol.SRoleList
}

func (this *SRoleList) Clone() MsgInfo {
	return new(SRoleList)
}

func (this *SRoleList) SetRoleId(r uint64) {
	this.RoleId = r
}

func (this *SRoleList) Unmarshal(data []byte) error {
	err := proto.Unmarshal(data, &this.Proto)
	return err
}

func (this *SRoleList) Process() {
	new(SRoleListProcess).Process(this)
}
