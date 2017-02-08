package message

import (
	"gameproject/common"
	"gameproject/world/msgProto"
	"gameproject/world/transMgr"

	"github.com/golang/protobuf/proto"
)

type W2SDispatch struct {
	msgProto.W2SDispatch
	s ISend  // Server缩写
}

func (this *W2SDispatch) Clone() MsgInfo {
	return new(W2SDispatch)
}

func (this *W2SDispatch) MsgType() uint32 {
	return 105
}

func (this *W2SDispatch) GetMsg() proto.Message {
	return &this.W2SDispatch
}

// 避免与协议的函数名称重复，以下函数命名有点特殊
func (this *W2SDispatch) Sets(s ISend) {
	this.s = s
}

func (this *W2SDispatch) Gets() ISend {
	return this.s
}

func (this *W2SDispatch) Unmarshal(data []byte) error {
	err := proto.Unmarshal(data, &this.W2SDispatch)
	return err
}

func (this *W2SDispatch) Send(msg MsgInfo) error {
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

func (this *W2SDispatch) Process(t *transMgr.Trans) bool {
	// do nothing
	return false
}
