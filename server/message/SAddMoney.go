package message

import (
	"gameproject/server/protocol"

	"github.com/golang/protobuf/proto"
)

type SAddMoney struct {
	RoleId uint64
	Proto  protocol.SAddMoney
}

func (this *SAddMoney) Clone() MsgInfo {
	return new(SAddMoney)
}

func (this *SAddMoney) SetRoleId(r uint64) {
	this.RoleId = r
}

func (this *SAddMoney) Unmarshal(data []byte) error {
	err := proto.Unmarshal(data, &this.Proto)
	return err
}

func (this *SAddMoney) Process() {
	new(SAddMoneyProcess).Process(this)
}
