package csmsg

import (
	"gameproject/common"
	"gameproject/server/client/csproto"
	"gameproject/server/client/msgMgr"

	"github.com/golang/protobuf/proto"
)

type CCreateRole struct {
	csproto.CCreateRole
	l msgMgr.ISend  // Link缩写
	g msgMgr.ISend  // Global缩写
	w msgMgr.ISend  // World缩写
	r uint64 // RoleId缩写
}

func (this *CCreateRole) Clone() msgMgr.MsgInfo {
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

func (this *CCreateRole) Setl(s msgMgr.ISend) {
	this.l = s
}

func (this *CCreateRole) Getl() msgMgr.ISend {
	return this.l
}

func (this *CCreateRole) Setg(s msgMgr.ISend) {
	this.g = s
}

func (this *CCreateRole) Getg() msgMgr.ISend {
	return this.g
}

func (this *CCreateRole) Setw(w msgMgr.ISend) {
	this.w = w
}

func (this *CCreateRole) Getw() msgMgr.ISend {
	return this.w
}

func (this *CCreateRole) Unmarshal(data []byte) error {
	err := proto.Unmarshal(data, &this.CCreateRole)
	return err
}

func (this *CCreateRole) Send2Link(msg msgMgr.MsgInfo) error {
	data, err := msgMgr.MarshalMsg(msg)
	if err != nil {
		return err
	}
	this.Getl().Send(data)
	return nil
}

func (this *CCreateRole) Process(t *common.Trans) bool {
	p := msgMgr.GetProc("CCreateRole")
	p.SetMsg(this)
	p.SetTrans(t)
	return p.Process()
}
