package csmsg

import (
	"gameproject/common"
	"gameproject/server/client/csproto"

	"github.com/golang/protobuf/proto"
)

type SEnterWorld struct {
	csproto.SEnterWorld
	l ISend  // Link缩写
	g ISend  // Global缩写
	w ISend  // World缩写
	r uint64 // RoleId缩写
}

func (this *SEnterWorld) Clone() MsgInfo {
	return new(SEnterWorld)
}

func (this *SEnterWorld) MsgType() uint32 {
	return 1008
}

func (this *SEnterWorld) GetMsg() proto.Message {
	return &this.SEnterWorld
}

// 避免与协议的函数名称重复，以下函数命名有点特殊
func (this *SEnterWorld) Setr(r uint64) {
	this.r = r
}

func (this *SEnterWorld) Getr() uint64 {
	return this.r
}

func (this *SEnterWorld) Setl(s ISend) {
	this.l = s
}

func (this *SEnterWorld) Getl() ISend {
	return this.l
}

func (this *SEnterWorld) Setg(s ISend) {
	this.g = s
}

func (this *SEnterWorld) Getg() ISend {
	return this.g
}

func (this *SEnterWorld) Setw(w ISend) {
	this.w = w
}

func (this *SEnterWorld) Getw() ISend {
	return this.w
}

func (this *SEnterWorld) Unmarshal(data []byte) error {
	err := proto.Unmarshal(data, &this.SEnterWorld)
	return err
}

func (this *SEnterWorld) Send2Link(msg MsgInfo) error {
	data, err := MarshalMsg(msg)
	if err != nil {
		return err
	}
	this.Getl().Send(data)
	return nil
}

func (this *SEnterWorld) Process(t *common.Trans) bool {
	// do nothing
	return false
}
