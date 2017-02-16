package csmsg

import (
	"gameproject/common"
	"gameproject/server/client/csproto"
	"gameproject/server/client/msgMgr"

	"github.com/golang/protobuf/proto"
)

type CAddLevel struct {
	csproto.CAddLevel
	l msgMgr.ISend  // Link缩写
	g msgMgr.ISend  // Global缩写
	w msgMgr.ISend  // World缩写
	r uint64 // RoleId缩写
}

func (this *CAddLevel) Clone() msgMgr.MsgInfo {
	return new(CAddLevel)
}

func (this *CAddLevel) MsgType() uint32 {
	return 1011
}

func (this *CAddLevel) GetMsg() proto.Message {
	return &this.CAddLevel
}

// 避免与协议的函数名称重复，以下函数命名有点特殊
func (this *CAddLevel) Setr(r uint64) {
	this.r = r
}

func (this *CAddLevel) Getr() uint64 {
	return this.r
}

func (this *CAddLevel) Setl(s msgMgr.ISend) {
	this.l = s
}

func (this *CAddLevel) Getl() msgMgr.ISend {
	return this.l
}

func (this *CAddLevel) Setg(s msgMgr.ISend) {
	this.g = s
}

func (this *CAddLevel) Getg() msgMgr.ISend {
	return this.g
}

func (this *CAddLevel) Setw(w msgMgr.ISend) {
	this.w = w
}

func (this *CAddLevel) Getw() msgMgr.ISend {
	return this.w
}

func (this *CAddLevel) Unmarshal(data []byte) error {
	err := proto.Unmarshal(data, &this.CAddLevel)
	return err
}

func (this *CAddLevel) Send2Link(msg msgMgr.MsgInfo) error {
	data, err := msgMgr.MarshalMsg(msg)
	if err != nil {
		return err
	}
	this.Getl().Send(data)
	return nil
}

func (this *CAddLevel) Process(t *common.Trans) bool {
	p := msgMgr.GetProc("CAddLevel")
	p.SetMsg(this)
	p.SetTrans(t)
	return p.Process()
}
