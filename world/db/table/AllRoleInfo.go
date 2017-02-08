package table

import (
	"gameproject/common"
	"gameproject/world/db/cacheMgr"
	"gameproject/world/db/dbProto"
	"gameproject/world/lockMgr"
	"gameproject/world/transMgr"
	"log"

	"github.com/golang/protobuf/proto"
)

type AllRoleInfo struct {
	dbProto.AllRoleInfo
	k string
}

func (this *AllRoleInfo) IsSave() bool {
	return true
}

func NewAllRoleInfo(t *transMgr.Trans, k string) *AllRoleInfo {
	r := new(AllRoleInfo)
	r.k = "AllRoleInfo_" + k
	if t != nil {
		t.Save(r)
	}
	return r
}

func GetAllRoleInfo(t *transMgr.Trans, k string) *AllRoleInfo {
	if k == "" {
		return nil
	}
	key := "AllRoleInfo_" + k
	t.Lock(key)
	v := cacheMgr.GetKV(key)
	if v == "" {
		return nil
	}
	
	oct := common.NewOctets([]byte(v))
	size := oct.UnmarshalUint32()
	if size != oct.Remain() {
		log.Panic("get table.AllRoleInfo Data Len Error:", key, ",", size, ",", len(v))
		return nil
	}
	data := oct.UnmarshalBytesOnly(size)
	r := NewAllRoleInfo(t, k)
	err := proto.Unmarshal(data, &r.AllRoleInfo)
	if err != nil {
		log.Panic("get DB Data Unmarshal Error:", r.k)
		return nil
	}
	return r
}

func SelectAllRoleInfo(k string) *AllRoleInfo {
	if k == "" {
		return nil
	}
	key := "AllRoleInfo_" + k
	lockMgr.Lock(key)
	defer lockMgr.Unlock(key)
	v := cacheMgr.GetKV(key)
	if v == "" {
		return nil
	}
	
	oct := common.NewOctets([]byte(v))
	size := oct.UnmarshalUint32()
	if size != oct.Remain() {
		log.Panic("select table.AllRoleInfo Data Len Error:", key, ",", size, ",", len(v))
		return nil
	}
	data := oct.UnmarshalBytesOnly(size)
	r := NewAllRoleInfo(nil, k)
	err := proto.Unmarshal(data, &r.AllRoleInfo)
	if err != nil {
		log.Panic("select DB Data Unmarshal Error:", r.k)
		return nil
	}
	return r
}

func (this *AllRoleInfo) Save() error {
	if this.k == "" {
		log.Panic("DB Data Save Error:", this.k)
	}
	data, err := proto.Marshal(&this.AllRoleInfo)
	if err != nil {
		return err
	}
	oct := &common.Octets{}
	oct.MarshalUint32(uint32(len(data)))
	oct.MarshalBytesOnly(data)
	cacheMgr.SetKV(this.k, string(oct.GetBuf()))
	return nil
}
