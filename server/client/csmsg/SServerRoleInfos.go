package csmsg

import (
	"gameproject/common"
	"gameproject/server/client/csproto"
	"gameproject/server/client/msgMgr"

	"github.com/golang/protobuf/proto"
)

type SServerRoleInfos struct {
	csproto.SServerRoleInfos
	l msgMgr.ISend  // Link缩写
	g msgMgr.ISend  // Global缩写
	w msgMgr.ISend  // World缩写
	r uint64 // RoleId缩写
}

func (this *SServerRoleInfos) Clone() msgMgr.MsgInfo {
	return new(SServerRoleInfos)
}

func (this *SServerRoleInfos) MsgType() uint32 {
	return 1014
}

func (this *SServerRoleInfos) GetMsg() proto.Message {
	return &this.SServerRoleInfos
}

// 避免与协议的函数名称重复，以下函数命名有点特殊
func (this *SServerRoleInfos) Setr(r uint64) {
	this.r = r
}

func (this *SServerRoleInfos) Getr() uint64 {
	return this.r
}

func (this *SServerRoleInfos) Setl(s msgMgr.ISend) {
	this.l = s
}

func (this *SServerRoleInfos) Getl() msgMgr.ISend {
	return this.l
}

func (this *SServerRoleInfos) Setg(s msgMgr.ISend) {
	this.g = s
}

func (this *SServerRoleInfos) Getg() msgMgr.ISend {
	return this.g
}

func (this *SServerRoleInfos) Setw(w msgMgr.ISend) {
	this.w = w
}

func (this *SServerRoleInfos) Getw() msgMgr.ISend {
	return this.w
}

func (this *SServerRoleInfos) Unmarshal(data []byte) error {
	err := proto.Unmarshal(data, &this.SServerRoleInfos)
	return err
}

func (this *SServerRoleInfos) Send2Link(msg msgMgr.MsgInfo) error {
	data, err := msgMgr.MarshalMsg(msg)
	if err != nil {
		return err
	}
	this.Getl().Send(data)
	return nil
}

func (this *SServerRoleInfos) Process(t *common.Trans) bool {
	// do nothing
	return false
}
