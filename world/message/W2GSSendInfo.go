package message

import (
	"gameproject/common"
	"gameproject/world/msgProto"
	"gameproject/world/transMgr"

	"github.com/golang/protobuf/proto"
)

type W2GSSendInfo struct {
	msgProto.W2GSSendInfo
	s ISend  // Server缩写
}

func (this *W2GSSendInfo) Clone() MsgInfo {
	return new(W2GSSendInfo)
}

func (this *W2GSSendInfo) MsgType() uint32 {
	return 104
}

func (this *W2GSSendInfo) GetMsg() proto.Message {
	return &this.W2GSSendInfo
}

// 避免与协议的函数名称重复，以下函数命名有点特殊
func (this *W2GSSendInfo) Sets(s ISend) {
	this.s = s
}

func (this *W2GSSendInfo) Gets() ISend {
	return this.s
}

func (this *W2GSSendInfo) Unmarshal(data []byte) error {
	err := proto.Unmarshal(data, &this.W2GSSendInfo)
	return err
}

func (this *W2GSSendInfo) Send(msg MsgInfo) error {
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

func (this *W2GSSendInfo) Process(t *transMgr.Trans) bool {
	// do nothing
	return false
}
