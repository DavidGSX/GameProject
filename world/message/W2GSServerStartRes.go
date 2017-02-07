package message

import (
	"gameproject/common"
	"gameproject/world/msgProto"
	"gameproject/world/transMgr"

	"github.com/golang/protobuf/proto"
)

type W2GSServerStartRes struct {
	msgProto.W2GSServerStartRes
	s ISend  // Server缩写
}

func (this *W2GSServerStartRes) Clone() MsgInfo {
	return new(W2GSServerStartRes)
}

func (this *W2GSServerStartRes) MsgType() uint32 {
	return 102
}

func (this *W2GSServerStartRes) GetMsg() proto.Message {
	return &this.W2GSServerStartRes
}

// 避免与协议的函数名称重复，以下函数命名有点特殊
func (this *W2GSServerStartRes) Sets(s ISend) {
	this.s = s
}

func (this *W2GSServerStartRes) Gets() ISend {
	return this.s
}

func (this *W2GSServerStartRes) Unmarshal(data []byte) error {
	err := proto.Unmarshal(data, &this.W2GSServerStartRes)
	return err
}

func (this *W2GSServerStartRes) Send(msg MsgInfo) error {
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

func (this *W2GSServerStartRes) Process(t *transMgr.Trans) bool {
	// do nothing
	return false
}
