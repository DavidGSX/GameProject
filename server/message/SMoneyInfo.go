package message

import (
	"gameproject/common"
	"gameproject/server/msgProto"
	"gameproject/server/transMgr"

	"github.com/golang/protobuf/proto"
)

type SMoneyInfo struct {
	msgProto.SMoneyInfo
	l ISend  // Link缩写
	g ISend  // Global缩写
	r uint64 // RoleId缩写
}

func (this *SMoneyInfo) Clone() MsgInfo {
	return new(SMoneyInfo)
}

func (this *SMoneyInfo) MsgType() uint32 {
	return 1008
}

func (this *SMoneyInfo) GetMsg() proto.Message {
	return &this.SMoneyInfo
}

// 避免与协议的函数名称重复，以下函数命名有点特殊
func (this *SMoneyInfo) Setr(r uint64) {
	this.r = r
}

func (this *SMoneyInfo) Getr() uint64 {
	return this.r
}

func (this *SMoneyInfo) Setl(s ISend) {
	this.l = s
}

func (this *SMoneyInfo) Getl() ISend {
	return this.l
}

func (this *SMoneyInfo) Setg(s ISend) {
	this.g = s
}

func (this *SMoneyInfo) Getg() ISend {
	return this.g
}

func (this *SMoneyInfo) Unmarshal(data []byte) error {
	err := proto.Unmarshal(data, &this.SMoneyInfo)
	return err
}

func (this *SMoneyInfo) Send(msg MsgInfo) error {
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

func (this *SMoneyInfo) Process(t *transMgr.Trans) bool {
	// do nothing
	return false
}
