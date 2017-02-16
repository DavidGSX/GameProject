package csmsg

import (
	"gameproject/common"
	"gameproject/server/client/csproto"
	"gameproject/server/client/msgMgr"

	"github.com/golang/protobuf/proto"
)

type CRoleList struct {
	csproto.CRoleList
	l msgMgr.ISend  // Link缩写
	g msgMgr.ISend  // Global缩写
	w msgMgr.ISend  // World缩写
	r uint64 // RoleId缩写
}

func (this *CRoleList) Clone() msgMgr.MsgInfo {
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

func (this *CRoleList) Setl(s msgMgr.ISend) {
	this.l = s
}

func (this *CRoleList) Getl() msgMgr.ISend {
	return this.l
}

func (this *CRoleList) Setg(s msgMgr.ISend) {
	this.g = s
}

func (this *CRoleList) Getg() msgMgr.ISend {
	return this.g
}

func (this *CRoleList) Setw(w msgMgr.ISend) {
	this.w = w
}

func (this *CRoleList) Getw() msgMgr.ISend {
	return this.w
}

func (this *CRoleList) Unmarshal(data []byte) error {
	err := proto.Unmarshal(data, &this.CRoleList)
	return err
}

func (this *CRoleList) Send2Link(msg msgMgr.MsgInfo) error {
	data, err := msgMgr.MarshalMsg(msg)
	if err != nil {
		return err
	}
	this.Getl().Send(data)
	return nil
}

func (this *CRoleList) Process(t *common.Trans) bool {
	p := msgMgr.GetProc("CRoleList")
	p.SetMsg(this)
	p.SetTrans(t)
	return p.Process()
}
