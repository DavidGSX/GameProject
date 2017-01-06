package message

import (
	"gameproject/server/protocol"

	"github.com/golang/protobuf/proto"
)

type SCreateRole struct {
	Link   ISend
	Global ISend
	RoleId uint64
	Proto  protocol.SCreateRole
}

func (this *SCreateRole) Clone() MsgInfo {
	return new(SCreateRole)
}

func (this *SCreateRole) SetRoleId(r uint64) {
	this.RoleId = r
}

func (this *SCreateRole) SetLink(s ISend) {
	this.Link = s
}

func (this *SCreateRole) SetGlobal(s ISend) {
	this.Global = s
}

func (this *SCreateRole) Unmarshal(data []byte) error {
	err := proto.Unmarshal(data, &this.Proto)
	return err
}

func (this *SCreateRole) Process() {
	//	new(SCreateRoleProcess).Process(this)
}
