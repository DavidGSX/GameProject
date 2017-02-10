package csmsg

import (
	"gameproject/common"
	"gameproject/server/client/csproto"

	"github.com/golang/protobuf/proto"
)

type CUserLogin struct {
	csproto.CUserLogin
	l ISend  // Link缩写
	g ISend  // Global缩写
	w ISend  // World缩写
	r uint64 // RoleId缩写
}

func (this *CUserLogin) Clone() MsgInfo {
	return new(CUserLogin)
}

func (this *CUserLogin) MsgType() uint32 {
	return 1001
}

func (this *CUserLogin) GetMsg() proto.Message {
	return &this.CUserLogin
}

// 避免与协议的函数名称重复，以下函数命名有点特殊
func (this *CUserLogin) Setr(r uint64) {
	this.r = r
}

func (this *CUserLogin) Getr() uint64 {
	return this.r
}

func (this *CUserLogin) Setl(s ISend) {
	this.l = s
}

func (this *CUserLogin) Getl() ISend {
	return this.l
}

func (this *CUserLogin) Setg(s ISend) {
	this.g = s
}

func (this *CUserLogin) Getg() ISend {
	return this.g
}

func (this *CUserLogin) Setw(w ISend) {
	this.w = w
}

func (this *CUserLogin) Getw() ISend {
	return this.w
}

func (this *CUserLogin) Unmarshal(data []byte) error {
	err := proto.Unmarshal(data, &this.CUserLogin)
	return err
}

func (this *CUserLogin) Send2Link(msg MsgInfo) error {
	data, err := MarshalMsg(msg)
	if err != nil {
		return err
	}
	this.Getl().Send(data)
	return nil
}

func (this *CUserLogin) Process(t *common.Trans) bool {
	p := new(CUserLoginProcess)
	p.msg = this
	p.trans = t
	return p.Process()
}
