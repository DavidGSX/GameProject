package message

import (
	"gameproject/server/protocol"

	"github.com/golang/protobuf/proto"
)

type CUserLogin struct {
	RoleId uint64
	Proto  protocol.CUserLogin
}

func (this *CUserLogin) Clone() MsgInfo {
	return new(CUserLogin)
}

func (this *CUserLogin) SetRoleId(r uint64) {
	this.RoleId = r
}

func (this *CUserLogin) Unmarshal(data []byte) error {
	err := proto.Unmarshal(data, &this.Proto)
	return err
}

func (this *CUserLogin) Process() {
	new(CUserLoginProcess).Process(this)
}
