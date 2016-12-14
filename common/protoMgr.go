package common

import (
	"log"
)

type IProtocal interface {
	Clone() IProtocal
	Init(*Linker)
	Process()
	Unmarshal(*Octets)
	Marshal() *Octets
}

var protoMgr *ProtoMgr

func GetProtoMgr() *ProtoMgr {
	if protoMgr == nil {
		protoMgr = new(ProtoMgr)
		protoMgr.init()
	}
	return protoMgr
}

type ProtoMgr struct {
	protos map[int]IProtocal
}

func (pMgr *ProtoMgr) init() {
	pMgr.protos = make(map[int]IProtocal)
}

func (pMgr *ProtoMgr) Add(t int, p IProtocal) {
	pMgr.protos[t] = p
}

func (pMgr *ProtoMgr) remove(t int) {
	delete(pMgr.protos, t)
}

func (pMgr *ProtoMgr) Process(t int, o *Octets, l *Linker) {
	proto, ok := pMgr.protos[t]
	if ok {
		pro := proto.Clone()
		pro.Unmarshal(o)
		pro.Init(l)
		pro.Process()
	} else {
		log.Panic("ProtoMgr Process Unkown Proto Type:", t)
	}
}
