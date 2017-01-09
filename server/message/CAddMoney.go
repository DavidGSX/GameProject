package message

import (
	"gameproject/common"
	"gameproject/server/protocol"

	"github.com/golang/protobuf/proto"
)

type CAddMoney struct {
	protocol.CAddMoney
	l ISend  // Link缩写
	g ISend  // Global缩写
	r uint64 // RoleId缩写
}

func (this *CAddMoney) Clone() MsgInfo {
	return new(CAddMoney)
}

func (this *CAddMoney) MsgType() uint32 {
	return 1007
}

// 避免与协议的函数名称重复，以下函数命名有点特殊
func (this *CAddMoney) Setr(r uint64) {
	this.r = r
}

func (this *CAddMoney) Getr() uint64 {
	return this.r
}

func (this *CAddMoney) Setl(s ISend) {
	this.l = s
}

func (this *CAddMoney) Getl() ISend {
	return this.l
}

func (this *CAddMoney) Setg(s ISend) {
	this.g = s
}

func (this *CAddMoney) Getg() ISend {
	return this.g
}

func (this *CAddMoney) Unmarshal(data []byte) error {
	err := proto.Unmarshal(data, &this.CAddMoney)
	return err
}

func (this *CAddMoney) Send(msg proto.Message) error {
	data, err := proto.Marshal(msg)
	if err != nil {
		return err
	}
	oct := &common.Octets{}
	oct.MarshalUint32(uint32(len(data)))
	oct.MarshalUint32(this.MsgType())
	oct.MarshalBytesOnly(data)
	this.Getl().Send(oct.GetBuf())
	return nil
}

func (this *CAddMoney) Process() {
	p := new(CAddMoneyProcess)
	p.CAddMoney = *this
	p.Process()
}
