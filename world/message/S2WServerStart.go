package message

import (
	"gameproject/common"
	"gameproject/world/msgProto"
	"gameproject/world/transMgr"

	"github.com/golang/protobuf/proto"
)

type S2WServerStart struct {
	msgProto.S2WServerStart
	s ISend  // Server缩写
}

func (this *S2WServerStart) Clone() MsgInfo {
	return new(S2WServerStart)
}

func (this *S2WServerStart) MsgType() uint32 {
	return 101
}

func (this *S2WServerStart) GetMsg() proto.Message {
	return &this.S2WServerStart
}

// 避免与协议的函数名称重复，以下函数命名有点特殊
func (this *S2WServerStart) Sets(s ISend) {
	this.s = s
}

func (this *S2WServerStart) Gets() ISend {
	return this.s
}

func (this *S2WServerStart) Unmarshal(data []byte) error {
	err := proto.Unmarshal(data, &this.S2WServerStart)
	return err
}

func (this *S2WServerStart) Send(msg MsgInfo) error {
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

func (this *S2WServerStart) Process(t *transMgr.Trans) bool {
	p := new(S2WServerStartProcess)
	p.msg = this
	p.trans = t
	return p.Process()
}
