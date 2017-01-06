package message

import (
	"gameproject/server/protocol"

	"github.com/golang/protobuf/proto"
)

type CCreateRole struct {
	Link   ISend
	Global ISend
	RoleId uint64
	Proto  protocol.CCreateRole
}

func (this *CCreateRole) Clone() MsgInfo {
	return new(CCreateRole)
}

func (this *CCreateRole) SetRoleId(r uint64) {
	this.RoleId = r
}

func (this *CCreateRole) SetLink(s ISend) {
	this.Link = s
}

func (this *CCreateRole) SetGlobal(s ISend) {
	this.Global = s
}

func (this *CCreateRole) Unmarshal(data []byte) error {
	err := proto.Unmarshal(data, &this.Proto)
	return err
}

func (this *CCreateRole) Process() {
	new(CCreateRoleProcess).Process(this)
}
