package message

import (
	"gameproject/common"
	"gameproject/world/msgProto"
	"gameproject/world/transMgr"

	"github.com/golang/protobuf/proto"
)

type GS2WSendInfo struct {
	msgProto.GS2WSendInfo
	s ISend  // Server缩写
}

func (this *GS2WSendInfo) Clone() MsgInfo {
	return new(GS2WSendInfo)
}

func (this *GS2WSendInfo) MsgType() uint32 {
	return 103
}

func (this *GS2WSendInfo) GetMsg() proto.Message {
	return &this.GS2WSendInfo
}

// 避免与协议的函数名称重复，以下函数命名有点特殊
func (this *GS2WSendInfo) Sets(s ISend) {
	this.s = s
}

func (this *GS2WSendInfo) Gets() ISend {
	return this.s
}

func (this *GS2WSendInfo) Unmarshal(data []byte) error {
	err := proto.Unmarshal(data, &this.GS2WSendInfo)
	return err
}

func (this *GS2WSendInfo) Send(msg MsgInfo) error {
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

func (this *GS2WSendInfo) Process(t *transMgr.Trans) bool {
	p := new(GS2WSendInfoProcess)
	p.msg = this
	p.trans = t
	return p.Process()
}
