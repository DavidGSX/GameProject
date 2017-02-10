package csmsg

import (
	"gameproject/common"
	"gameproject/server/client/csproto"

	"github.com/golang/protobuf/proto"
)

type SLevelInfo struct {
	csproto.SLevelInfo
	l ISend  // Link缩写
	g ISend  // Global缩写
	w ISend  // World缩写
	r uint64 // RoleId缩写
}

func (this *SLevelInfo) Clone() MsgInfo {
	return new(SLevelInfo)
}

func (this *SLevelInfo) MsgType() uint32 {
	return 1012
}

func (this *SLevelInfo) GetMsg() proto.Message {
	return &this.SLevelInfo
}

// 避免与协议的函数名称重复，以下函数命名有点特殊
func (this *SLevelInfo) Setr(r uint64) {
	this.r = r
}

func (this *SLevelInfo) Getr() uint64 {
	return this.r
}

func (this *SLevelInfo) Setl(s ISend) {
	this.l = s
}

func (this *SLevelInfo) Getl() ISend {
	return this.l
}

func (this *SLevelInfo) Setg(s ISend) {
	this.g = s
}

func (this *SLevelInfo) Getg() ISend {
	return this.g
}

func (this *SLevelInfo) Setw(w ISend) {
	this.w = w
}

func (this *SLevelInfo) Getw() ISend {
	return this.w
}

func (this *SLevelInfo) Unmarshal(data []byte) error {
	err := proto.Unmarshal(data, &this.SLevelInfo)
	return err
}

func (this *SLevelInfo) Send2Link(msg MsgInfo) error {
	data, err := MarshalMsg(msg)
	if err != nil {
		return err
	}
	this.Getl().Send(data)
	return nil
}

func (this *SLevelInfo) Process(t *common.Trans) bool {
	// do nothing
	return false
}
