package message

import (
	"gameproject/server/protocol"

	"github.com/golang/protobuf/proto"
)

type SMoneyInfo struct {
	Link   ISend
	Global ISend
	RoleId uint64
	Proto  protocol.SMoneyInfo
}

func (this *SMoneyInfo) Clone() MsgInfo {
	return new(SMoneyInfo)
}

func (this *SMoneyInfo) SetRoleId(r uint64) {
	this.RoleId = r
}

func (this *SMoneyInfo) SetLink(s ISend) {
	this.Link = s
}

func (this *SMoneyInfo) SetGlobal(s ISend) {
	this.Global = s
}

func (this *SMoneyInfo) Unmarshal(data []byte) error {
	err := proto.Unmarshal(data, &this.Proto)
	return err
}

func (this *SMoneyInfo) Process() {
	//	new(SMoneyInfoProcess).Process(this)
}
