package message

import (
	"gameproject/common"
	"gameproject/server/msgProto"

	"github.com/golang/protobuf/proto"
)

type SCreateRole struct {
	msgProto.SCreateRole
	l ISend  // Link缩写
	g ISend  // Global缩写
	r uint64 // RoleId缩写
}

func (this *SCreateRole) Clone() MsgInfo {
	return new(SCreateRole)
}

func (this *SCreateRole) MsgType() uint32 {
	return 1006
}

func (this *SCreateRole) GetMsg() proto.Message {
	return &this.SCreateRole
}

// 避免与协议的函数名称重复，以下函数命名有点特殊
func (this *SCreateRole) Setr(r uint64) {
	this.r = r
}

func (this *SCreateRole) Getr() uint64 {
	return this.r
}

func (this *SCreateRole) Setl(s ISend) {
	this.l = s
}

func (this *SCreateRole) Getl() ISend {
	return this.l
}

func (this *SCreateRole) Setg(s ISend) {
	this.g = s
}

func (this *SCreateRole) Getg() ISend {
	return this.g
}

func (this *SCreateRole) Unmarshal(data []byte) error {
	err := proto.Unmarshal(data, &this.SCreateRole)
	return err
}

func (this *SCreateRole) Send(msg MsgInfo) error {
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

func (this *SCreateRole) Process() {
	// do nothing
}
