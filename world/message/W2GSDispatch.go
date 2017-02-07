package message

import (
	"gameproject/common"
	"gameproject/world/msgProto"
	"gameproject/world/transMgr"

	"github.com/golang/protobuf/proto"
)

type W2GSDispatch struct {
	msgProto.W2GSDispatch
	s ISend  // Server缩写
}

func (this *W2GSDispatch) Clone() MsgInfo {
	return new(W2GSDispatch)
}

func (this *W2GSDispatch) MsgType() uint32 {
	return 105
}

func (this *W2GSDispatch) GetMsg() proto.Message {
	return &this.W2GSDispatch
}

// 避免与协议的函数名称重复，以下函数命名有点特殊
func (this *W2GSDispatch) Sets(s ISend) {
	this.s = s
}

func (this *W2GSDispatch) Gets() ISend {
	return this.s
}

func (this *W2GSDispatch) Unmarshal(data []byte) error {
	err := proto.Unmarshal(data, &this.W2GSDispatch)
	return err
}

func (this *W2GSDispatch) Send(msg MsgInfo) error {
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

func (this *W2GSDispatch) Process(t *transMgr.Trans) bool {
	// do nothing
	return false
}
