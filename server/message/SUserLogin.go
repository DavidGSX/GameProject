package message

import (
	"gameproject/server/protocol"

	"github.com/golang/protobuf/proto"
)

type SUserLogin struct {
	Link   ISend
	Global ISend
	RoleId uint64
	Proto  protocol.SUserLogin
}

func (this *SUserLogin) Clone() MsgInfo {
	return new(SUserLogin)
}

func (this *SUserLogin) SetRoleId(r uint64) {
	this.RoleId = r
}

func (this *SUserLogin) SetLink(s ISend) {
	this.Link = s
}

func (this *SUserLogin) SetGlobal(s ISend) {
	this.Global = s
}

func (this *SUserLogin) Unmarshal(data []byte) error {
	err := proto.Unmarshal(data, &this.Proto)
	return err
}

func (this *SUserLogin) Process() {
	new(SUserLoginProcess).Process(this)
}
