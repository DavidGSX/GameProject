package message

import (
	"gameproject/server/protocol"

	"github.com/golang/protobuf/proto"
)

type CRoleList struct {
	Link   ISend
	Global ISend
	RoleId uint64
	Proto  protocol.CRoleList
}

func (this *CRoleList) Clone() MsgInfo {
	return new(CRoleList)
}

func (this *CRoleList) SetRoleId(r uint64) {
	this.RoleId = r
}

func (this *CRoleList) SetLink(s ISend) {
	this.Link = s
}

func (this *CRoleList) SetGlobal(s ISend) {
	this.Global = s
}

func (this *CRoleList) Unmarshal(data []byte) error {
	err := proto.Unmarshal(data, &this.Proto)
	return err
}

func (this *CRoleList) Process() {
	new(CRoleListProcess).Process(this)
}
