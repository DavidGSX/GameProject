package message

import (
	"gameproject/common"
	"gameproject/world/msgProto"

	"github.com/golang/protobuf/proto"
)

type S2WSendInfo struct {
	msgProto.S2WSendInfo
	s ISend  // Server缩写
}

func (this *S2WSendInfo) Clone() MsgInfo {
	return new(S2WSendInfo)
}

func (this *S2WSendInfo) MsgType() uint32 {
	return 103
}

func (this *S2WSendInfo) GetMsg() proto.Message {
	return &this.S2WSendInfo
}

// 避免与协议的函数名称重复，以下函数命名有点特殊
func (this *S2WSendInfo) Sets(s ISend) {
	this.s = s
}

func (this *S2WSendInfo) Gets() ISend {
	return this.s
}

func (this *S2WSendInfo) Unmarshal(data []byte) error {
	err := proto.Unmarshal(data, &this.S2WSendInfo)
	return err
}

func (this *S2WSendInfo) Send(msg MsgInfo) error {
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

func (this *S2WSendInfo) Process(t *common.Trans) bool {
	p := new(S2WSendInfoProcess)
	p.msg = this
	p.trans = t
	return p.Process()
}
