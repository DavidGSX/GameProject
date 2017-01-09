package table

import (
	"gameproject/common"
	"gameproject/server/cacheMgr"
	"gameproject/server/dbProto"
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

func NewProperty(k uint64) *Property {
	ret := new(Property)
	ret.k = "Property_" + strconv.FormatUint(k,10)
	return ret
}

func GetProperty(uk uint64) *Property {
	if uk == 0 {
		return nil
	}
	k := strconv.FormatUint(uk,10)
	v := cacheMgr.GetKV("Property_" + k)
	if v == "" {
		return nil
	}
	
	oct := common.NewOctets([]byte(v))
	size := oct.UnmarshalUint32()
	if size != oct.Remain() {
		log.Panic("table.Property Data Len Error:", k, ",", size, ",", len(v))
		return nil
	}
	data := oct.UnmarshalBytesOnly(size)
	t := NewProperty(uk)
	err := proto.Unmarshal(data, &t.Property)
	if err != nil {
		log.Panic("DB Data Unmarshal Error:", t.k)
		return nil
	}
	return t
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
