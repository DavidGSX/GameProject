package csmsg

import (
	"gameproject/common"
	"gameproject/server/client/csproto"
	"gameproject/server/client/msgMgr"

	"github.com/golang/protobuf/proto"
)

type CEnterWorld struct {
	csproto.CEnterWorld
	l msgMgr.ISend  // Link缩写
	g msgMgr.ISend  // Global缩写
	w msgMgr.ISend  // World缩写
	r uint64 // RoleId缩写
}

func (this *CEnterWorld) Clone() msgMgr.MsgInfo {
	return new(CEnterWorld)
}

func (this *CEnterWorld) MsgType() uint32 {
	return 1007
}

func (this *CEnterWorld) GetMsg() proto.Message {
	return &this.CEnterWorld
}

// 避免与协议的函数名称重复，以下函数命名有点特殊
func (this *CEnterWorld) Setr(r uint64) {
	this.r = r
}

func (this *CEnterWorld) Getr() uint64 {
	return this.r
}

func (this *CEnterWorld) Setl(s msgMgr.ISend) {
	this.l = s
}

func (this *CEnterWorld) Getl() msgMgr.ISend {
	return this.l
}

func (this *CEnterWorld) Setg(s msgMgr.ISend) {
	this.g = s
}

func (this *CEnterWorld) Getg() msgMgr.ISend {
	return this.g
}

func (this *CEnterWorld) Setw(w msgMgr.ISend) {
	this.w = w
}

func (this *CEnterWorld) Getw() msgMgr.ISend {
	return this.w
}

func (this *CEnterWorld) Unmarshal(data []byte) error {
	err := proto.Unmarshal(data, &this.CEnterWorld)
	return err
}

func (this *CEnterWorld) Send2Link(msg msgMgr.MsgInfo) error {
	data, err := msgMgr.MarshalMsg(msg)
	if err != nil {
		return err
	}
	this.Getl().Send(data)
	return nil
}

func (this *CEnterWorld) Process(t *common.Trans) bool {
	p := msgMgr.GetProc("CEnterWorld")
	p.SetMsg(this)
	p.SetTrans(t)
	return p.Process()
}
