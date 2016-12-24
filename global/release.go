package global

import (
	"gameproject/common"
)

type Release struct {
	GroupName
}

func (p *Release) GetType() int {
	return 4
}

func (p *Release) Clone() common.IProtocal {
	return new(Release)
}

func (p *Release) Process() {
	r := NewResult(p.l)
	/*
		conn, err := common.GetDBPool().Get()
		if err != nil {
			log.Panic("Release Process Conn Get Error", err)
		}
		defer common.GetDBPool().Put(conn)

		_, err = conn.Cmd("HDEL", p.group, p.name).Int()
		if err != nil {
			log.Panic("Release Process Cmd Error", err)
		}

		r.SetRes(RPC_OK)
	*/
	r.Process()
}
