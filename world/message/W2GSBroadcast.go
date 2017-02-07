package message

import (
	"gameproject/common"
	"gameproject/world/msgProto"
	"gameproject/world/transMgr"

	"github.com/golang/protobuf/proto"
)

type W2GSBroadcast struct {
	msgProto.W2GSBroadcast
	s ISend  // Server缩写
}

func (this *W2GSBroadcast) Clone() MsgInfo {
	return new(W2GSBroadcast)
}

func (this *W2GSBroadcast) MsgType() uint32 {
	return 106
}

func (this *W2GSBroadcast) GetMsg() proto.Message {
	return &this.W2GSBroadcast
}

// 避免与协议的函数名称重复，以下函数命名有点特殊
func (this *W2GSBroadcast) Sets(s ISend) {
	this.s = s
}

func (this *W2GSBroadcast) Gets() ISend {
	return this.s
}

func (this *W2GSBroadcast) Unmarshal(data []byte) error {
	err := proto.Unmarshal(data, &this.W2GSBroadcast)
	return err
}

func (this *W2GSBroadcast) Send(msg MsgInfo) error {
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

func (this *W2GSBroadcast) Process(t *transMgr.Trans) bool {
	// do nothing
	return false
}
