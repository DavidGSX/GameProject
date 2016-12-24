package global

import (
	"gameproject/common"
)

const (
	RPC_OK         = 12345
	RPC_DUPLICATE  = 1
	RPC_NOT_EXISTS = 2
)

type Result struct {
	l   *common.Linker
	res int
}

func (p *Result) GetType() int {
	return 1
}

func (p *Result) Clone() *Result {
	return new(Result)
}

func NewResult(l *common.Linker) *Result {
	r := new(Result)
	r.l = l
	return r
}

func (p *Result) Process() {
	o := new(common.Octets)
	o.CompactUint32(uint32(p.GetType()))
	o.MarshalBytes(p.Marshal().GetBuf())
	p.l.Send(o.GetBuf())
}

func (p *Result) SetRes(r int) {
	p.res = r
}

func (p *Result) Unmarshal(o *common.Octets) {
	p.res = int(o.UnmarshalUint32())
}

func (p *Result) Marshal() *common.Octets {
	o := new(common.Octets)
	o.MarshalUint32(uint32(p.l.GetSid())) //sid
	o.MarshalUint32(uint32(p.res))        //result
	return o
}
