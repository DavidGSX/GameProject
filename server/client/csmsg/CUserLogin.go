package csmsg

import (
	"gameproject/common"
	"gameproject/server/client/csproto"
	"gameproject/server/client/msgMgr"

	"github.com/golang/protobuf/proto"
)

type CUserLogin struct {
	csproto.CUserLogin
	l msgMgr.ISend  // Link缩写
	g msgMgr.ISend  // Global缩写
	w msgMgr.ISend  // World缩写
	r uint64 // RoleId缩写
}

func (this *CUserLogin) Clone() msgMgr.MsgInfo {
	return new(CUserLogin)
}

func (this *CUserLogin) MsgType() uint32 {
	return 1001
}

func (this *CUserLogin) GetMsg() proto.Message {
	return &this.CUserLogin
}

// 避免与协议的函数名称重复，以下函数命名有点特殊
func (this *CUserLogin) Setr(r uint64) {
	this.r = r
}

func (this *CUserLogin) Getr() uint64 {
	return this.r
}

func (this *CUserLogin) Setl(s msgMgr.ISend) {
	this.l = s
}

func (this *CUserLogin) Getl() msgMgr.ISend {
	return this.l
}

func (this *CUserLogin) Setg(s msgMgr.ISend) {
	this.g = s
}

func (this *CUserLogin) Getg() msgMgr.ISend {
	return this.g
}

func (this *CUserLogin) Setw(w msgMgr.ISend) {
	this.w = w
}

func (this *CUserLogin) Getw() msgMgr.ISend {
	return this.w
}

func (this *CUserLogin) Unmarshal(data []byte) error {
	err := proto.Unmarshal(data, &this.CUserLogin)
	return err
}

func (this *CUserLogin) Send2Link(msg msgMgr.MsgInfo) error {
	data, err := msgMgr.MarshalMsg(msg)
	if err != nil {
		return err
	}
	this.Getl().Send(data)
	return nil
}

func (this *CUserLogin) Process(t *common.Trans) bool {
	p := msgMgr.GetProc("CUserLogin")
	p.SetMsg(this)
	p.SetTrans(t)
	return p.Process()
}
