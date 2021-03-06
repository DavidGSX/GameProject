package csmsg

import (
	"gameproject/common"
	"gameproject/server/client/csproto"
	"gameproject/server/client/msgMgr"

	"github.com/golang/protobuf/proto"
)

type SRoleList struct {
	csproto.SRoleList
	l msgMgr.ISend  // Link缩写
	g msgMgr.ISend  // Global缩写
	w msgMgr.ISend  // World缩写
	r uint64 // RoleId缩写
}

func (this *SRoleList) Clone() msgMgr.MsgInfo {
	return new(SRoleList)
}

func (this *SRoleList) MsgType() uint32 {
	return 1004
}

func (this *SRoleList) GetMsg() proto.Message {
	return &this.SRoleList
}

// 避免与协议的函数名称重复，以下函数命名有点特殊
func (this *SRoleList) Setr(r uint64) {
	this.r = r
}

func (this *SRoleList) Getr() uint64 {
	return this.r
}

func (this *SRoleList) Setl(s msgMgr.ISend) {
	this.l = s
}

func (this *SRoleList) Getl() msgMgr.ISend {
	return this.l
}

func (this *SRoleList) Setg(s msgMgr.ISend) {
	this.g = s
}

func (this *SRoleList) Getg() msgMgr.ISend {
	return this.g
}

func (this *SRoleList) Setw(w msgMgr.ISend) {
	this.w = w
}

func (this *SRoleList) Getw() msgMgr.ISend {
	return this.w
}

func (this *SRoleList) Unmarshal(data []byte) error {
	err := proto.Unmarshal(data, &this.SRoleList)
	return err
}

func (this *SRoleList) Send2Link(msg msgMgr.MsgInfo) error {
	data, err := msgMgr.MarshalMsg(msg)
	if err != nil {
		return err
	}
	this.Getl().Send(data)
	return nil
}

func (this *SRoleList) Process(t *common.Trans) bool {
	// do nothing
	return false
}
