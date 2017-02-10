package csmsg

import (
	"gameproject/common"
	"gameproject/server/client/csproto"

	"github.com/golang/protobuf/proto"
)

type SUserLogin struct {
	csproto.SUserLogin
	l ISend  // Link缩写
	g ISend  // Global缩写
	w ISend  // World缩写
	r uint64 // RoleId缩写
}

func (this *SUserLogin) Clone() MsgInfo {
	return new(SUserLogin)
}

func (this *SUserLogin) MsgType() uint32 {
	return 1002
}

func (this *SUserLogin) GetMsg() proto.Message {
	return &this.SUserLogin
}

// 避免与协议的函数名称重复，以下函数命名有点特殊
func (this *SUserLogin) Setr(r uint64) {
	this.r = r
}

func (this *SUserLogin) Getr() uint64 {
	return this.r
}

func (this *SUserLogin) Setl(s ISend) {
	this.l = s
}

func (this *SUserLogin) Getl() ISend {
	return this.l
}

func (this *SUserLogin) Setg(s ISend) {
	this.g = s
}

func (this *SUserLogin) Getg() ISend {
	return this.g
}

func (this *SUserLogin) Setw(w ISend) {
	this.w = w
}

func (this *SUserLogin) Getw() ISend {
	return this.w
}

func (this *SUserLogin) Unmarshal(data []byte) error {
	err := proto.Unmarshal(data, &this.SUserLogin)
	return err
}

func (this *SUserLogin) Send2Link(msg MsgInfo) error {
	data, err := MarshalMsg(msg)
	if err != nil {
		return err
	}
	this.Getl().Send(data)
	return nil
}

func (this *SUserLogin) Process(t *common.Trans) bool {
	// do nothing
	return false
}
