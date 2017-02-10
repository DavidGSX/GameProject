package csmsg

import (
	"gameproject/common"
	"gameproject/server/client/csproto"

	"github.com/golang/protobuf/proto"
)

type CEnterWorld struct {
	csproto.CEnterWorld
	l ISend  // Link缩写
	g ISend  // Global缩写
	w ISend  // World缩写
	r uint64 // RoleId缩写
}

func (this *CEnterWorld) Clone() MsgInfo {
	return new(CEnterWorld)
}

func (this *CEnterWorld) MsgType() uint32 {
	return 1007
}

func (this *CEnterWorld) GetMsg() proto.Message {
	return &this.CEnterWorld
}

// 避免与协议的函数名称重复，以下函数命名有点特殊
func (this *CEnterWorld) Setr(r uint64) {
	this.r = r
}

func (this *CEnterWorld) Getr() uint64 {
	return this.r
}

func (this *CEnterWorld) Setl(s ISend) {
	this.l = s
}

func (this *CEnterWorld) Getl() ISend {
	return this.l
}

func (this *CEnterWorld) Setg(s ISend) {
	this.g = s
}

func (this *CEnterWorld) Getg() ISend {
	return this.g
}

func (this *CEnterWorld) Setw(w ISend) {
	this.w = w
}

func (this *CEnterWorld) Getw() ISend {
	return this.w
}

func (this *CEnterWorld) Unmarshal(data []byte) error {
	err := proto.Unmarshal(data, &this.CEnterWorld)
	return err
}

func (this *CEnterWorld) Send2Link(msg MsgInfo) error {
	data, err := MarshalMsg(msg)
	if err != nil {
		return err
	}
	this.Getl().Send(data)
	return nil
}

func (this *CEnterWorld) Process(t *common.Trans) bool {
	p := new(CEnterWorldProcess)
	p.msg = this
	p.trans = t
	return p.Process()
}
