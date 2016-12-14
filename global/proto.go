package global

import (
	"gameproject/common"
	"log"
	"unicode/utf16"
)

type GroupName struct {
	group   string
	name    string
	localid int
	reserve *common.Octets
	l       *common.Linker
}

func (p *GroupName) Unmarshal(o *common.Octets) {
	p.group = string(utf16.Decode(o.UnmarshalUint16s())) //o.UnmarshalString()
	p.name = string(utf16.Decode(o.UnmarshalUint16s()))
	p.localid = int(o.UnmarshalUint32())
	p.reserve = common.NewOctets(o.UnmarshalBytes())
	log.Println("group:", p.group, "name:", p.name, "localid:", p.localid, "reserve:", p.reserve.GetBuf())
}

func (p *GroupName) Marshal() *common.Octets {
	o := new(common.Octets)
	o.MarshalString(p.group)
	o.MarshalUint16s(utf16.Encode([]rune(p.name)))
	o.MarshalUint32(uint32(p.localid))
	o.MarshalBytes(p.reserve.GetBuf())
	return o
}

func (p *GroupName) Init(l *common.Linker) {
	p.l = l
}
