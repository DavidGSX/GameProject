package message

import (
	"gameproject/common"
	"gameproject/world/msgProto"
	"gameproject/world/transMgr"

	"github.com/golang/protobuf/proto"
)

type W2SRoleInfoResponse struct {
	msgProto.W2SRoleInfoResponse
	s ISend  // Server缩写
}

func (this *W2SRoleInfoResponse) Clone() MsgInfo {
	return new(W2SRoleInfoResponse)
}

func (this *W2SRoleInfoResponse) MsgType() uint32 {
	return 109
}

func (this *W2SRoleInfoResponse) GetMsg() proto.Message {
	return &this.W2SRoleInfoResponse
}

// 避免与协议的函数名称重复，以下函数命名有点特殊
func (this *W2SRoleInfoResponse) Sets(s ISend) {
	this.s = s
}

func (this *W2SRoleInfoResponse) Gets() ISend {
	return this.s
}

func (this *W2SRoleInfoResponse) Unmarshal(data []byte) error {
	err := proto.Unmarshal(data, &this.W2SRoleInfoResponse)
	return err
}

func (this *W2SRoleInfoResponse) Send(msg MsgInfo) error {
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

func (this *W2SRoleInfoResponse) Process(t *transMgr.Trans) bool {
	// do nothing
	return false
}
