package csmsg

import (
	"gameproject/common"
	"gameproject/server/client/csproto"
	"gameproject/server/client/msgMgr"

	"github.com/golang/protobuf/proto"
)

type CReqServerRoleInfos struct {
	csproto.CReqServerRoleInfos
	l msgMgr.ISend  // Link缩写
	g msgMgr.ISend  // Global缩写
	w msgMgr.ISend  // World缩写
	r uint64 // RoleId缩写
}

func (this *CReqServerRoleInfos) Clone() msgMgr.MsgInfo {
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

func (this *CReqServerRoleInfos) Setl(s msgMgr.ISend) {
	this.l = s
}

func (this *CReqServerRoleInfos) Getl() msgMgr.ISend {
	return this.l
}

func (this *CReqServerRoleInfos) Setg(s msgMgr.ISend) {
	this.g = s
}

func (this *CReqServerRoleInfos) Getg() msgMgr.ISend {
	return this.g
}

func (this *CReqServerRoleInfos) Setw(w msgMgr.ISend) {
	this.w = w
}

func (this *CReqServerRoleInfos) Getw() msgMgr.ISend {
	return this.w
}

func (this *CReqServerRoleInfos) Unmarshal(data []byte) error {
	err := proto.Unmarshal(data, &this.CReqServerRoleInfos)
	return err
}

func (this *CReqServerRoleInfos) Send2Link(msg msgMgr.MsgInfo) error {
	data, err := msgMgr.MarshalMsg(msg)
	if err != nil {
		return err
	}
	this.Getl().Send(data)
	return nil
}

func (this *CReqServerRoleInfos) Process(t *common.Trans) bool {
	p := msgMgr.GetProc("CReqServerRoleInfos")
	p.SetMsg(this)
	p.SetTrans(t)
	return p.Process()
}
