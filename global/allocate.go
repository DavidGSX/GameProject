package global

import (
	"gameproject/common"
)

type Allocate struct {
	GroupName
}

func (p *Allocate) GetType() int {
	return 2
}

func (p *Allocate) Clone() common.IProtocal {
	return new(Allocate)
}

func (p *Allocate) Process() {
	r := NewResult(p.l)
	/*
		conn, err := common.GetDBPool().Get()
		if err != nil {
			log.Panic("Allocate Process Conn Get Error", err)
		}
		defer common.GetDBPool().Put(conn)

		resp, err := conn.Cmd("HSET", p.group, p.name, "1").Int()
		if err != nil {
			log.Panic("Allocate Process Cmd Error ", err)
		}

		if resp == 1 {
			r.SetRes(RPC_OK)
		} else {
			r.SetRes(RPC_DUPLICATE)
		}
	*/
	r.Process()
}
