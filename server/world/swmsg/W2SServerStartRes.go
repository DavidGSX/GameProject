package swmsg

import (
	"gameproject/common"
	"gameproject/world/msgProto"

	"github.com/golang/protobuf/proto"
)

type W2SServerStartRes struct {
	msgProto.W2SServerStartRes
	s ISend  // Server缩写
}

func (this *W2SServerStartRes) Clone() MsgInfo {
	return new(W2SServerStartRes)
}

func (this *W2SServerStartRes) MsgType() uint32 {
	return 102
}

func (this *W2SServerStartRes) GetMsg() proto.Message {
	return &this.W2SServerStartRes
}

// 避免与协议的函数名称重复，以下函数命名有点特殊
func (this *W2SServerStartRes) Sets(s ISend) {
	this.s = s
}

func (this *W2SServerStartRes) Gets() ISend {
	return this.s
}

func (this *W2SServerStartRes) Unmarshal(data []byte) error {
	err := proto.Unmarshal(data, &this.W2SServerStartRes)
	return err
}

func (this *W2SServerStartRes) Send(msg MsgInfo) error {
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

func (this *W2SServerStartRes) Process(t *common.Trans) bool {
	p := new(W2SServerStartResProcess)
	p.msg = this
	p.trans = t
	return p.Process()
}
