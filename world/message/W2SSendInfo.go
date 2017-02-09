package message

import (
	"gameproject/common"
	"gameproject/world/msgProto"

	"github.com/golang/protobuf/proto"
)

type W2SSendInfo struct {
	msgProto.W2SSendInfo
	s ISend  // Server缩写
}

func (this *W2SSendInfo) Clone() MsgInfo {
	return new(W2SSendInfo)
}

func (this *W2SSendInfo) MsgType() uint32 {
	return 104
}

func (this *W2SSendInfo) GetMsg() proto.Message {
	return &this.W2SSendInfo
}

// 避免与协议的函数名称重复，以下函数命名有点特殊
func (this *W2SSendInfo) Sets(s ISend) {
	this.s = s
}

func (this *W2SSendInfo) Gets() ISend {
	return this.s
}

func (this *W2SSendInfo) Unmarshal(data []byte) error {
	err := proto.Unmarshal(data, &this.W2SSendInfo)
	return err
}

func (this *W2SSendInfo) Send(msg MsgInfo) error {
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

func (this *W2SSendInfo) Process(t *common.Trans) bool {
	// do nothing
	return false
}
