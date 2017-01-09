package table

import (
	"gameproject/common"
	"gameproject/server/cacheMgr"
	"gameproject/server/dbProto"
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

func NewUser(k string) *User {
	ret := new(User)
	ret.k = "User_" + k
	return ret
}

func GetUser(k string) *User {
	if k == "" {
		return nil
	}
	v := cacheMgr.GetKV("User_" + k)
	if v == "" {
		return nil
	}
	
	oct := common.NewOctets([]byte(v))
	size := oct.UnmarshalUint32()
	if size != oct.Remain() {
		log.Panic("table.User Data Len Error:", k, ",", size, ",", len(v))
		return nil
	}
	data := oct.UnmarshalBytesOnly(size)
	t := NewUser(k)
	err := proto.Unmarshal(data, &t.User)
	if err != nil {
		log.Panic("DB Data Unmarshal Error:", t.k)
		return nil
	}
	return t
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
