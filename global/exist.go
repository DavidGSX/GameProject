package global

import (
	"gameproject/common"
	"log"
)

type Exist struct {
	GroupName
}

func (p *Exist) GetType() int {
	return 1
}

func (p *Exist) Clone() common.IProtocal {
	return new(Exist)
}

func (p *Exist) Process() {
	r := NewResult(p.l)

	conn, err := common.GetDBPool().Get()
	if err != nil {
		log.Panic("Exist Process Conn Get Error", err)
	}
	defer common.GetDBPool().Put(conn)

	resp := conn.Cmd("HGET", p.group, p.name)
	if resp.Err != nil {
		log.Panic("Exist Process Cmd Error ", err)
	}

	res := resp.String()
	if res == "Resp(Nil)" {
		r.SetRes(RPC_NOT_EXISTS)
	} else if res == "Resp(1)" {
		r.SetRes(RPC_OK)
	}

	r.Process()
}
