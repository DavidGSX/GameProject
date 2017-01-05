package message

import (
	"gameproject/server/protocol"

	"github.com/golang/protobuf/proto"
)

type CUserLogin struct {
	Link   ISend
	Global ISend
	RoleId uint64
	Proto  protocol.CUserLogin
}

func (this *CUserLogin) Clone() MsgInfo {
	return new(CUserLogin)
}

func (this *CUserLogin) SetRoleId(r uint64) {
	this.RoleId = r
}

func (this *CUserLogin) SetLink(s ISend) {
	this.Link = s
}

func (this *CUserLogin) SetGlobal(s ISend) {
	this.Global = s
}

func (this *CUserLogin) Unmarshal(data []byte) error {
	err := proto.Unmarshal(data, &this.Proto)
	return err
}

func (this *CUserLogin) Process() {
	new(CUserLoginProcess).Process(this)
}
