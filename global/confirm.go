package global

import (
	"gameproject/common"
)

type Confirm struct {
	GroupName
}

func (p *Confirm) GetType() int {
	return 3
}

func (p *Confirm) Clone() common.IProtocal {
	return new(Confirm)
}

func (p *Confirm) Process() {
	r := NewResult(p.l)
	r.SetRes(RPC_OK)
	r.Process()
}
