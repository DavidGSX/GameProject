package message

import (
	"gameproject/common"
	"gameproject/global/protocol"
	"log"

	"github.com/golang/protobuf/proto"
)

type CUserLoginProcess struct {
	CUserLogin
}

func (this *CUserLoginProcess) Process() {
	defer func() {
		if err := recover(); err != nil {
			log.Println("CUserLoginProcess Error:", err)
		}
	}()

	userId := this.UserId
	token := this.Token

	log.Println("UserId:", userId, " Token:", token)
	this.Getl().SetUserId(userId)

	send := &protocol.SGUserAuth{}
	send.UserId = userId
	send.Token = token
	data, err := proto.Marshal(send)
	if err != nil {
		log.Println("CUserLoginProcess Marshal SGUserAuth error:", err)
		return
	}
	oct := &common.Octets{}
	oct.MarshalUint32(uint32(len(data)))
	oct.MarshalUint32(2)
	oct.MarshalBytesOnly(data)
	this.Getg().Send(oct.GetBuf())
}
