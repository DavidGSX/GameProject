package csmsg

import (
	"gameproject/common"
	"gameproject/server/client/csproto"

	"github.com/golang/protobuf/proto"
)

type CAddLevel struct {
	csproto.CAddLevel
	l ISend  // Link缩写
	g ISend  // Global缩写
	w ISend  // World缩写
	r uint64 // RoleId缩写
}

func (this *CAddLevel) Clone() MsgInfo {
	return new(CAddLevel)
}

func (this *CAddLevel) MsgType() uint32 {
	return 1011
}

func (this *CAddLevel) GetMsg() proto.Message {
	return &this.CAddLevel
}

// 避免与协议的函数名称重复，以下函数命名有点特殊
func (this *CAddLevel) Setr(r uint64) {
	this.r = r
}

func (this *CAddLevel) Getr() uint64 {
	return this.r
}

func (this *CAddLevel) Setl(s ISend) {
	this.l = s
}

func (this *CAddLevel) Getl() ISend {
	return this.l
}

func (this *CAddLevel) Setg(s ISend) {
	this.g = s
}

func (this *CAddLevel) Getg() ISend {
	return this.g
}

func (this *CAddLevel) Setw(w ISend) {
	this.w = w
}

func (this *CAddLevel) Getw() ISend {
	return this.w
}

func (this *CAddLevel) Unmarshal(data []byte) error {
	err := proto.Unmarshal(data, &this.CAddLevel)
	return err
}

func (this *CAddLevel) Send2Link(msg MsgInfo) error {
	data, err := MarshalMsg(msg)
	if err != nil {
		return err
	}
	this.Getl().Send(data)
	return nil
}

func (this *CAddLevel) Process(t *common.Trans) bool {
	p := new(CAddLevelProcess)
	p.msg = this
	p.trans = t
	return p.Process()
}
