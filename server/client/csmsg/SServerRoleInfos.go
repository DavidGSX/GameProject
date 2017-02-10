package csmsg

import (
	"gameproject/common"
	"gameproject/server/client/csproto"

	"github.com/golang/protobuf/proto"
)

type SServerRoleInfos struct {
	csproto.SServerRoleInfos
	l ISend  // Link缩写
	g ISend  // Global缩写
	w ISend  // World缩写
	r uint64 // RoleId缩写
}

func (this *SServerRoleInfos) Clone() MsgInfo {
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

func (this *SServerRoleInfos) Setl(s ISend) {
	this.l = s
}

func (this *SServerRoleInfos) Getl() ISend {
	return this.l
}

func (this *SServerRoleInfos) Setg(s ISend) {
	this.g = s
}

func (this *SServerRoleInfos) Getg() ISend {
	return this.g
}

func (this *SServerRoleInfos) Setw(w ISend) {
	this.w = w
}

func (this *SServerRoleInfos) Getw() ISend {
	return this.w
}

func (this *SServerRoleInfos) Unmarshal(data []byte) error {
	err := proto.Unmarshal(data, &this.SServerRoleInfos)
	return err
}

func (this *SServerRoleInfos) Send2Link(msg MsgInfo) error {
	data, err := MarshalMsg(msg)
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
