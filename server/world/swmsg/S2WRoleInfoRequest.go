package swmsg

import (
	"gameproject/common"
	"gameproject/world/msgProto"

	"github.com/golang/protobuf/proto"
)

type S2WRoleInfoRequest struct {
	msgProto.S2WRoleInfoRequest
	s ISend  // Server缩写
}

func (this *S2WRoleInfoRequest) Clone() MsgInfo {
	return new(S2WRoleInfoRequest)
}

func (this *S2WRoleInfoRequest) MsgType() uint32 {
	return 108
}

func (this *S2WRoleInfoRequest) GetMsg() proto.Message {
	return &this.S2WRoleInfoRequest
}

// 避免与协议的函数名称重复，以下函数命名有点特殊
func (this *S2WRoleInfoRequest) Sets(s ISend) {
	this.s = s
}

func (this *S2WRoleInfoRequest) Gets() ISend {
	return this.s
}

func (this *S2WRoleInfoRequest) Unmarshal(data []byte) error {
	err := proto.Unmarshal(data, &this.S2WRoleInfoRequest)
	return err
}

func (this *S2WRoleInfoRequest) Send(msg MsgInfo) error {
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

func (this *S2WRoleInfoRequest) Process(t *common.Trans) bool {
	// do nothing
	return false
}
