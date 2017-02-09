package table

import (
	"gameproject/common"
	"gameproject/common/cache"
	"gameproject/server/db/dbProto"
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

func NewProperty(t *common.Trans, k uint64) *Property {
	r := new(Property)
	r.k = "Property_" + strconv.FormatUint(k, 10)
	if t != nil {
		t.Save(r)
	}
	return r
}

func GetProperty(t *common.Trans, k uint64) *Property {
	if k == 0 {
		return nil
	}
	key := "Property_" + strconv.FormatUint(k, 10)
	t.Lock(key)
	v := cache.GetKV(key)
	if v == "" {
		return nil
	}
	
	oct := common.NewOctets([]byte(v))
	size := oct.UnmarshalUint32()
	if size != oct.Remain() {
		log.Panic("get table.Property Data Len Error:", key, ",", size, ",", len(v))
		return nil
	}
	data := oct.UnmarshalBytesOnly(size)
	r := NewProperty(t, k)
	err := proto.Unmarshal(data, &r.Property)
	if err != nil {
		log.Panic("get DB Data Unmarshal Error:", r.k)
		return nil
	}
	return r
}

func SelectProperty(k uint64) *Property {
	if k == 0 {
		return nil
	}
	key := "Property_" + strconv.FormatUint(k, 10)
	common.Lock(key)
	defer common.Unlock(key)
	v := cache.GetKV(key)
	if v == "" {
		return nil
	}
	
	oct := common.NewOctets([]byte(v))
	size := oct.UnmarshalUint32()
	if size != oct.Remain() {
		log.Panic("select table.Property Data Len Error:", key, ",", size, ",", len(v))
		return nil
	}
	data := oct.UnmarshalBytesOnly(size)
	r := NewProperty(nil, k)
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
	cache.SetKV(this.k, string(oct.GetBuf()))
	return nil
}
