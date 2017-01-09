package table

import (
	"gameproject/common"
	"gameproject/server/cacheMgr"
	"gameproject/server/dbProto"
	"log"

	"github.com/golang/protobuf/proto"
)

type Name struct {
	dbProto.Name
	k string
}

func (this *Name) IsSave() bool {
	return true
}

func NewName(k string) *Name {
	ret := new(Name)
	ret.k = "Name_" + k
	return ret
}

func GetName(k string) *Name {
	if k == "" {
		return nil
	}
	v := cacheMgr.GetKV("Name_" + k)
	if v == "" {
		return nil
	}
	
	oct := common.NewOctets([]byte(v))
	size := oct.UnmarshalUint32()
	if size != oct.Remain() {
		log.Panic("table.Name Data Len Error:", k, ",", size, ",", len(v))
		return nil
	}
	data := oct.UnmarshalBytesOnly(size)
	t := NewName(k)
	err := proto.Unmarshal(data, &t.Name)
	if err != nil {
		log.Panic("DB Data Unmarshal Error:", t.k)
		return nil
	}
	return t
}

func (this *Name) Save() error {
	if this.k == "" {
		log.Panic("DB Data Save Error:", this.k)
	}
	data, err := proto.Marshal(&this.Name)
	if err != nil {
		return err
	}
	oct := &common.Octets{}
	oct.MarshalUint32(uint32(len(data)))
	oct.MarshalBytesOnly(data)
	cacheMgr.SetKV(this.k, string(oct.GetBuf()))
	return nil
}
