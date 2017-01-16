package message

import (
	"gameproject/common"
	"gameproject/server/msgProto"
	"gameproject/server/transMgr"

	"github.com/golang/protobuf/proto"
)

type SUserLogin struct {
	msgProto.SUserLogin
	l ISend  // Link缩写
	g ISend  // Global缩写
	r uint64 // RoleId缩写
}

func (this *SUserLogin) Clone() MsgInfo {
	return new(SUserLogin)
}

func (this *SUserLogin) MsgType() uint32 {
	return 1002
}

func (this *SUserLogin) GetMsg() proto.Message {
	return &this.SUserLogin
}

// 避免与协议的函数名称重复，以下函数命名有点特殊
func (this *SUserLogin) Setr(r uint64) {
	this.r = r
}

func (this *SUserLogin) Getr() uint64 {
	return this.r
}

func (this *SUserLogin) Setl(s ISend) {
	this.l = s
}

func (this *SUserLogin) Getl() ISend {
	return this.l
}

func (this *SUserLogin) Setg(s ISend) {
	this.g = s
}

func (this *SUserLogin) Getg() ISend {
	return this.g
}

func (this *SUserLogin) Unmarshal(data []byte) error {
	err := proto.Unmarshal(data, &this.SUserLogin)
	return err
}

func (this *SUserLogin) Send(msg MsgInfo) error {
	data, err := proto.Marshal(msg.GetMsg())
	if err != nil {
		return err
	}
	oct := &common.Octets{}
	oct.MarshalUint32(uint32(len(data)))
	oct.MarshalUint32(msg.MsgType())
	oct.MarshalBytesOnly(data)
	this.Getl().Send(oct.GetBuf())
	return nil
}

func (this *SUserLogin) Process(t *transMgr.Trans) bool {
	// do nothing
	return false
}
