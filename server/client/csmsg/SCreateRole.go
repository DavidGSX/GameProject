package csmsg

import (
	"gameproject/common"
	"gameproject/server/client/csproto"
	"gameproject/server/client/msgMgr"

	"github.com/golang/protobuf/proto"
)

type SCreateRole struct {
	csproto.SCreateRole
	l msgMgr.ISend  // Link缩写
	g msgMgr.ISend  // Global缩写
	w msgMgr.ISend  // World缩写
	r uint64 // RoleId缩写
}

func (this *SCreateRole) Clone() msgMgr.MsgInfo {
	return new(SCreateRole)
}

func (this *SCreateRole) MsgType() uint32 {
	return 1006
}

func (this *SCreateRole) GetMsg() proto.Message {
	return &this.SCreateRole
}

// 避免与协议的函数名称重复，以下函数命名有点特殊
func (this *SCreateRole) Setr(r uint64) {
	this.r = r
}

func (this *SCreateRole) Getr() uint64 {
	return this.r
}

func (this *SCreateRole) Setl(s msgMgr.ISend) {
	this.l = s
}

func (this *SCreateRole) Getl() msgMgr.ISend {
	return this.l
}

func (this *SCreateRole) Setg(s msgMgr.ISend) {
	this.g = s
}

func (this *SCreateRole) Getg() msgMgr.ISend {
	return this.g
}

func (this *SCreateRole) Setw(w msgMgr.ISend) {
	this.w = w
}

func (this *SCreateRole) Getw() msgMgr.ISend {
	return this.w
}

func (this *SCreateRole) Unmarshal(data []byte) error {
	err := proto.Unmarshal(data, &this.SCreateRole)
	return err
}

func (this *SCreateRole) Send2Link(msg msgMgr.MsgInfo) error {
	data, err := msgMgr.MarshalMsg(msg)
	if err != nil {
		return err
	}
	this.Getl().Send(data)
	return nil
}

func (this *SCreateRole) Process(t *common.Trans) bool {
	// do nothing
	return false
}
