package message

import (
	"gameproject/common"
	"gameproject/server/msgProto"

	"github.com/golang/protobuf/proto"
)

type CRoleList struct {
	msgProto.CRoleList
	l ISend  // Link缩写
	g ISend  // Global缩写
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

func (this *CRoleList) Unmarshal(data []byte) error {
	err := proto.Unmarshal(data, &this.CRoleList)
	return err
}

func (this *CRoleList) Send(msg MsgInfo) error {
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

func (this *CRoleList) Process() {
	p := new(CRoleListProcess)
	p.CRoleList = *this
	p.Process()
}
