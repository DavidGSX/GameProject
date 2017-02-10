package csmsg

import (
	"gameproject/common"
	"gameproject/server/client/csproto"

	"github.com/golang/protobuf/proto"
)

type CRoleList struct {
	csproto.CRoleList
	l ISend  // Link缩写
	g ISend  // Global缩写
	w ISend  // World缩写
	r uint64 // RoleId缩写
}

func (this *CRoleList) Clone() MsgInfo {
	return new(CRoleList)
}

func (this *CRoleList) MsgType() uint32 {
	return 1003
}

func (this *CRoleList) GetMsg() proto.Message {
	return &this.CRoleList
}

// 避免与协议的函数名称重复，以下函数命名有点特殊
func (this *CRoleList) Setr(r uint64) {
	this.r = r
}

func (this *CRoleList) Getr() uint64 {
	return this.r
}

func (this *CRoleList) Setl(s ISend) {
	this.l = s
}

func (this *CRoleList) Getl() ISend {
	return this.l
}

func (this *CRoleList) Setg(s ISend) {
	this.g = s
}

func (this *CRoleList) Getg() ISend {
	return this.g
}

func (this *CRoleList) Setw(w ISend) {
	this.w = w
}

func (this *CRoleList) Getw() ISend {
	return this.w
}

func (this *CRoleList) Unmarshal(data []byte) error {
	err := proto.Unmarshal(data, &this.CRoleList)
	return err
}

func (this *CRoleList) Send2Link(msg MsgInfo) error {
	data, err := MarshalMsg(msg)
	if err != nil {
		return err
	}
	this.Getl().Send(data)
	return nil
}

func (this *CRoleList) Process(t *common.Trans) bool {
	p := new(CRoleListProcess)
	p.msg = this
	p.trans = t
	return p.Process()
}
