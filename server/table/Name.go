package table

import (
	"gameproject/common"
	"gameproject/server/cacheMgr"
	"gameproject/server/dbProto"
	"gameproject/server/lockMgr"
	"gameproject/server/transMgr"
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

func NewName(t *transMgr.Trans, k string) *Name {
	r := new(Name)
	r.k = "Name_" + k
	if t != nil {
		t.Save(r)
	}
	return r
}

func GetName(t *transMgr.Trans, k string) *Name {
	if k == "" {
		return nil
	}
	t.Lock("Name_" + k)
	v := cacheMgr.GetKV("Name_" + k)
	if v == "" {
		return nil
	}
	
	oct := common.NewOctets([]byte(v))
	size := oct.UnmarshalUint32()
	if size != oct.Remain() {
		log.Panic("get table.Name Data Len Error:", k, ",", size, ",", len(v))
		return nil
	}
	data := oct.UnmarshalBytesOnly(size)
	r := NewName(t, k)
	err := proto.Unmarshal(data, &r.Name)
	if err != nil {
		log.Panic("get DB Data Unmarshal Error:", r.k)
		return nil
	}
	return r
}

func SelectName(k string) *Name {
	if k == "" {
		return nil
	}
	lockMgr.Lock("Name_" + k)
	defer lockMgr.Unlock("Name_" + k)
	v := cacheMgr.GetKV("Name_" + k)
	if v == "" {
		return nil
	}
	
	oct := common.NewOctets([]byte(v))
	size := oct.UnmarshalUint32()
	if size != oct.Remain() {
		log.Panic("select table.Name Data Len Error:", k, ",", size, ",", len(v))
		return nil
	}
	data := oct.UnmarshalBytesOnly(size)
	r := NewName(nil, k)
	err := proto.Unmarshal(data, &r.Name)
	if err != nil {
		log.Panic("select DB Data Unmarshal Error:", r.k)
		return nil
	}
	return r
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
