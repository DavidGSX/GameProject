package csmsg

import (
	"gameproject/common"
	"gameproject/server/client/csproto"

	"github.com/golang/protobuf/proto"
)

type CReqServerRoleInfos struct {
	csproto.CReqServerRoleInfos
	l ISend  // Link缩写
	g ISend  // Global缩写
	w ISend  // World缩写
	r uint64 // RoleId缩写
}

func (this *CReqServerRoleInfos) Clone() MsgInfo {
	return new(CReqServerRoleInfos)
}

func (this *CReqServerRoleInfos) MsgType() uint32 {
	return 1013
}

func (this *CReqServerRoleInfos) GetMsg() proto.Message {
	return &this.CReqServerRoleInfos
}

// 避免与协议的函数名称重复，以下函数命名有点特殊
func (this *CReqServerRoleInfos) Setr(r uint64) {
	this.r = r
}

func (this *CReqServerRoleInfos) Getr() uint64 {
	return this.r
}

func (this *CReqServerRoleInfos) Setl(s ISend) {
	this.l = s
}

func (this *CReqServerRoleInfos) Getl() ISend {
	return this.l
}

func (this *CReqServerRoleInfos) Setg(s ISend) {
	this.g = s
}

func (this *CReqServerRoleInfos) Getg() ISend {
	return this.g
}

func (this *CReqServerRoleInfos) Setw(w ISend) {
	this.w = w
}

func (this *CReqServerRoleInfos) Getw() ISend {
	return this.w
}

func (this *CReqServerRoleInfos) Unmarshal(data []byte) error {
	err := proto.Unmarshal(data, &this.CReqServerRoleInfos)
	return err
}

func (this *CReqServerRoleInfos) Send2Link(msg MsgInfo) error {
	data, err := MarshalMsg(msg)
	if err != nil {
		return err
	}
	this.Getl().Send(data)
	return nil
}

func (this *CReqServerRoleInfos) Process(t *common.Trans) bool {
	p := new(CReqServerRoleInfosProcess)
	p.msg = this
	p.trans = t
	return p.Process()
}
