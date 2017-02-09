package message

import (
	"gameproject/common"
	"gameproject/world/msgProto"

	"github.com/golang/protobuf/proto"
)

type S2WRoleInfoChange struct {
	msgProto.S2WRoleInfoChange
	s ISend  // Server缩写
}

func (this *S2WRoleInfoChange) Clone() MsgInfo {
	return new(S2WRoleInfoChange)
}

func (this *S2WRoleInfoChange) MsgType() uint32 {
	return 107
}

func (this *S2WRoleInfoChange) GetMsg() proto.Message {
	return &this.S2WRoleInfoChange
}

// 避免与协议的函数名称重复，以下函数命名有点特殊
func (this *S2WRoleInfoChange) Sets(s ISend) {
	this.s = s
}

func (this *S2WRoleInfoChange) Gets() ISend {
	return this.s
}

func (this *S2WRoleInfoChange) Unmarshal(data []byte) error {
	err := proto.Unmarshal(data, &this.S2WRoleInfoChange)
	return err
}

func (this *S2WRoleInfoChange) Send(msg MsgInfo) error {
	data, err := proto.Marshal(msg.GetMsg())
	if err != nil {
		return err
	}
	oct := &common.Octets{}
	oct.MarshalUint32(uint32(len(data)))
	oct.MarshalUint32(msg.MsgType())
	oct.MarshalBytesOnly(data)
	this.Gets().Send(oct.GetBuf())
	return nil
}

func (this *S2WRoleInfoChange) Process(t *common.Trans) bool {
	p := new(S2WRoleInfoChangeProcess)
	p.msg = this
	p.trans = t
	return p.Process()
}
