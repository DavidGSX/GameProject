package message

import (
	"gameproject/server/protocol"

	"github.com/golang/protobuf/proto"
)

type CAddMoney struct {
	Link   ISend
	Global ISend
	RoleId uint64
	Proto  protocol.CAddMoney
}

func (this *CAddMoney) Clone() MsgInfo {
	return new(CAddMoney)
}

func (this *CAddMoney) SetRoleId(r uint64) {
	this.RoleId = r
}

func (this *CAddMoney) SetLink(s ISend) {
	this.Link = s
}

func (this *CAddMoney) SetGlobal(s ISend) {
	this.Global = s
}

func (this *CAddMoney) Unmarshal(data []byte) error {
	err := proto.Unmarshal(data, &this.Proto)
	return err
}

func (this *CAddMoney) Process() {
	new(CAddMoneyProcess).Process(this)
}
