package csmsg

import (
	"gameproject/common"
	"gameproject/server/client/csproto"

	"github.com/golang/protobuf/proto"
)

type CCreateRole struct {
	csproto.CCreateRole
	l ISend  // Link缩写
	g ISend  // Global缩写
	w ISend  // World缩写
	r uint64 // RoleId缩写
}

func (this *CCreateRole) Clone() MsgInfo {
	return new(CCreateRole)
}

func (this *CCreateRole) MsgType() uint32 {
	return 1005
}

func (this *CCreateRole) GetMsg() proto.Message {
	return &this.CCreateRole
}

// 避免与协议的函数名称重复，以下函数命名有点特殊
func (this *CCreateRole) Setr(r uint64) {
	this.r = r
}

func (this *CCreateRole) Getr() uint64 {
	return this.r
}

func (this *CCreateRole) Setl(s ISend) {
	this.l = s
}

func (this *CCreateRole) Getl() ISend {
	return this.l
}

func (this *CCreateRole) Setg(s ISend) {
	this.g = s
}

func (this *CCreateRole) Getg() ISend {
	return this.g
}

func (this *CCreateRole) Setw(w ISend) {
	this.w = w
}

func (this *CCreateRole) Getw() ISend {
	return this.w
}

func (this *CCreateRole) Unmarshal(data []byte) error {
	err := proto.Unmarshal(data, &this.CCreateRole)
	return err
}

func (this *CCreateRole) Send2Link(msg MsgInfo) error {
	data, err := MarshalMsg(msg)
	if err != nil {
		return err
	}
	this.Getl().Send(data)
	return nil
}

func (this *CCreateRole) Process(t *common.Trans) bool {
	p := new(CCreateRoleProcess)
	p.msg = this
	p.trans = t
	return p.Process()
}