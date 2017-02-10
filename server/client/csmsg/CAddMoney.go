package csmsg

import (
	"gameproject/common"
	"gameproject/server/client/csproto"

	"github.com/golang/protobuf/proto"
)

type CAddMoney struct {
	csproto.CAddMoney
	l ISend  // Link缩写
	g ISend  // Global缩写
	w ISend  // World缩写
	r uint64 // RoleId缩写
}

func (this *CAddMoney) Clone() MsgInfo {
	return new(CAddMoney)
}

func (this *CAddMoney) MsgType() uint32 {
	return 1009
}

func (this *CAddMoney) GetMsg() proto.Message {
	return &this.CAddMoney
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

func (this *CAddMoney) Setw(w ISend) {
	this.w = w
}

func (this *CAddMoney) Getw() ISend {
	return this.w
}

func (this *CAddMoney) Unmarshal(data []byte) error {
	err := proto.Unmarshal(data, &this.CAddMoney)
	return err
}

func (this *CAddMoney) Send2Link(msg MsgInfo) error {
	data, err := MarshalMsg(msg)
	if err != nil {
		return err
	}
	this.Getl().Send(data)
	return nil
}

func (this *CAddMoney) Process(t *common.Trans) bool {
	p := new(CAddMoneyProcess)
	p.msg = this
	p.trans = t
	return p.Process()
}
