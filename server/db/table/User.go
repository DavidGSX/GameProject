package table

import (
	"gameproject/common"
	"gameproject/server/db/cacheMgr"
	"gameproject/server/db/dbProto"
	"log"

	"github.com/golang/protobuf/proto"
)

type User struct {
	dbProto.User
	k string
}

func (this *User) IsSave() bool {
	return true
}

func NewUser(t *common.Trans, k string) *User {
	r := new(User)
	r.k = "User_" + k
	if t != nil {
		t.Save(r)
	}
	return r
}

func GetUser(t *common.Trans, k string) *User {
	if k == "" {
		return nil
	}
	key := "User_" + k
	t.Lock(key)
	v := cacheMgr.GetKV(key)
	if v == "" {
		return nil
	}
	
	oct := common.NewOctets([]byte(v))
	size := oct.UnmarshalUint32()
	if size != oct.Remain() {
		log.Panic("get table.User Data Len Error:", key, ",", size, ",", len(v))
		return nil
	}
	data := oct.UnmarshalBytesOnly(size)
	r := NewUser(t, k)
	err := proto.Unmarshal(data, &r.User)
	if err != nil {
		log.Panic("get DB Data Unmarshal Error:", r.k)
		return nil
	}
	return r
}

func SelectUser(k string) *User {
	if k == "" {
		return nil
	}
	key := "User_" + k
	common.Lock(key)
	defer common.Unlock(key)
	v := cacheMgr.GetKV(key)
	if v == "" {
		return nil
	}
	
	oct := common.NewOctets([]byte(v))
	size := oct.UnmarshalUint32()
	if size != oct.Remain() {
		log.Panic("select table.User Data Len Error:", key, ",", size, ",", len(v))
		return nil
	}
	data := oct.UnmarshalBytesOnly(size)
	r := NewUser(nil, k)
	err := proto.Unmarshal(data, &r.User)
	if err != nil {
		log.Panic("select DB Data Unmarshal Error:", r.k)
		return nil
	}
	return r
}

func (this *User) Save() error {
	if this.k == "" {
		log.Panic("DB Data Save Error:", this.k)
	}
	data, err := proto.Marshal(&this.User)
	if err != nil {
		return err
	}
	oct := &common.Octets{}
	oct.MarshalUint32(uint32(len(data)))
	oct.MarshalBytesOnly(data)
	cacheMgr.SetKV(this.k, string(oct.GetBuf()))
	return nil
}
