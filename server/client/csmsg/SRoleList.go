package csmsg

import (
	"gameproject/common"
	"gameproject/server/client/csproto"

	"github.com/golang/protobuf/proto"
)

type SRoleList struct {
	csproto.SRoleList
	l ISend  // Link缩写
	g ISend  // Global缩写
	r uint64 // RoleId缩写
}

func (this *SRoleList) Clone() MsgInfo {
	return new(SRoleList)
}

func (this *SRoleList) MsgType() uint32 {
	return 1004
}

func (this *SRoleList) GetMsg() proto.Message {
	return &this.SRoleList
}

// 避免与协议的函数名称重复，以下函数命名有点特殊
func (this *SRoleList) Setr(r uint64) {
	this.r = r
}

func (this *SRoleList) Getr() uint64 {
	return this.r
}

func (this *SRoleList) Setl(s ISend) {
	this.l = s
}

func (this *SRoleList) Getl() ISend {
	return this.l
}

func (this *SRoleList) Setg(s ISend) {
	this.g = s
}

func (this *SRoleList) Getg() ISend {
	return this.g
}

func (this *SRoleList) Unmarshal(data []byte) error {
	err := proto.Unmarshal(data, &this.SRoleList)
	return err
}

func (this *SRoleList) Send(msg MsgInfo) error {
	data, err := proto.Marshal(msg.GetMsg())
	if err != nil {
		return err
	}
	oct := &common.Octets{}
	oct.MarshalUint32(uint32(len(data)))
	oct.MarshalUint32(msg.MsgType())
	oct.MarshalBytesOnly(data)
	this.Getl().Send(oct.GetBuf())
	return nil
}

func (this *SRoleList) Process(t *common.Trans) bool {
	// do nothing
	return false
}
