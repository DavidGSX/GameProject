package table

import (
	"gameproject/common"
	"gameproject/server/cacheMgr"
	"gameproject/server/dbProto"
	"gameproject/server/lockMgr"
	"gameproject/server/transMgr"
	"log"
	"strconv"

	"github.com/golang/protobuf/proto"
)

type Property struct {
	dbProto.Property
	k string
}

func (this *Property) IsSave() bool {
	return true
}

func NewProperty(t *transMgr.Trans, k uint64) *Property {
	r := new(Property)
	r.k = "Property_" + strconv.FormatUint(k,10)
	if t != nil {
		t.Save(r)
	}
	return r
}

func GetProperty(t *transMgr.Trans, uk uint64) *Property {
	if uk == 0 {
		return nil
	}
	k := strconv.FormatUint(uk,10)
	t.Lock("Property_" + k)
	v := cacheMgr.GetKV("Property_" + k)
	if v == "" {
		return nil
	}
	
	oct := common.NewOctets([]byte(v))
	size := oct.UnmarshalUint32()
	if size != oct.Remain() {
		log.Panic("get table.Property Data Len Error:", k, ",", size, ",", len(v))
		return nil
	}
	data := oct.UnmarshalBytesOnly(size)
	r := NewProperty(t, uk)
	err := proto.Unmarshal(data, &r.Property)
	if err != nil {
		log.Panic("get DB Data Unmarshal Error:", r.k)
		return nil
	}
	return r
}

func SelectProperty(uk uint64) *Property {
	if uk == 0 {
		return nil
	}
	k := strconv.FormatUint(uk,10)
	lockMgr.Lock("Property_" + k)
	defer lockMgr.Unlock("Property_" + k)
	v := cacheMgr.GetKV("Property_" + k)
	if v == "" {
		return nil
	}
	
	oct := common.NewOctets([]byte(v))
	size := oct.UnmarshalUint32()
	if size != oct.Remain() {
		log.Panic("select table.Property Data Len Error:", k, ",", size, ",", len(v))
		return nil
	}
	data := oct.UnmarshalBytesOnly(size)
	r := NewProperty(nil, uk)
	err := proto.Unmarshal(data, &r.Property)
	if err != nil {
		log.Panic("select DB Data Unmarshal Error:", r.k)
		return nil
	}
	return r
}

func (this *Property) Save() error {
	if this.k == "" {
		log.Panic("DB Data Save Error:", this.k)
	}
	data, err := proto.Marshal(&this.Property)
	if err != nil {
		return err
	}
	oct := &common.Octets{}
	oct.MarshalUint32(uint32(len(data)))
	oct.MarshalBytesOnly(data)
	cacheMgr.SetKV(this.k, string(oct.GetBuf()))
	return nil
}
