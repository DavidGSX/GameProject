package message

import (
	"gameproject/common"
	"gameproject/server/protocol"

	"github.com/golang/protobuf/proto"
)

type CUserLogin struct {
	protocol.CUserLogin
	l ISend  // Link缩写
	g ISend  // Global缩写
	r uint64 // RoleId缩写
}

func (this *CUserLogin) Clone() MsgInfo {
	return new(CUserLogin)
}

func (this *CUserLogin) MsgType() uint32 {
	return 1001
}

// 避免与协议的函数名称重复，以下函数命名有点特殊
func (this *CUserLogin) Setr(r uint64) {
	this.r = r
}

func (this *CUserLogin) Getr() uint64 {
	return this.r
}

func (this *CUserLogin) Setl(s ISend) {
	this.l = s
}

func (this *CUserLogin) Getl() ISend {
	return this.l
}

func (this *CUserLogin) Setg(s ISend) {
	this.g = s
}

func (this *CUserLogin) Getg() ISend {
	return this.g
}

func (this *CUserLogin) Unmarshal(data []byte) error {
	err := proto.Unmarshal(data, &this.CUserLogin)
	return err
}

func (this *CUserLogin) Send(msg proto.Message) error {
	data, err := proto.Marshal(msg)
	if err != nil {
		return err
	}
	oct := &common.Octets{}
	oct.MarshalUint32(uint32(len(data)))
	oct.MarshalUint32(this.MsgType())
	oct.MarshalBytesOnly(data)
	this.Getl().Send(oct.GetBuf())
	return nil
}

func (this *CUserLogin) Process() {
	p := new(CUserLoginProcess)
	p.CUserLogin = *this
	p.Process()
}
