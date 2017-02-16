package csmsg

import (
	"gameproject/common"
	"gameproject/server/client/csproto"
	"gameproject/server/client/msgMgr"

	"github.com/golang/protobuf/proto"
)

type SLevelInfo struct {
	csproto.SLevelInfo
	l msgMgr.ISend  // Link缩写
	g msgMgr.ISend  // Global缩写
	w msgMgr.ISend  // World缩写
	r uint64 // RoleId缩写
}

func (this *SLevelInfo) Clone() msgMgr.MsgInfo {
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

func (this *SLevelInfo) Setl(s msgMgr.ISend) {
	this.l = s
}

func (this *SLevelInfo) Getl() msgMgr.ISend {
	return this.l
}

func (this *SLevelInfo) Setg(s msgMgr.ISend) {
	this.g = s
}

func (this *SLevelInfo) Getg() msgMgr.ISend {
	return this.g
}

func (this *SLevelInfo) Setw(w msgMgr.ISend) {
	this.w = w
}

func (this *SLevelInfo) Getw() msgMgr.ISend {
	return this.w
}

func (this *SLevelInfo) Unmarshal(data []byte) error {
	err := proto.Unmarshal(data, &this.SLevelInfo)
	return err
}

func (this *SLevelInfo) Send2Link(msg msgMgr.MsgInfo) error {
	data, err := msgMgr.MarshalMsg(msg)
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
