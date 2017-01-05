package message

import (
	"gameproject/common"
	"gameproject/global/protocol"
	"log"

	"github.com/golang/protobuf/proto"
)

type CUserLoginProcess struct {
	msg *CUserLogin
}

func (this *CUserLoginProcess) Process(msg *CUserLogin) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("CUserLoginProcess Error:", err)
		}
	}()

	userId := msg.Proto.UserId
	token := msg.Proto.Token

	log.Println("UserId:", userId, " Token:", token)
	msg.Link.SetUserId(userId)

	send := &protocol.SGUserAuth{}
	send.UserId = userId
	send.Token = token
	data, err := proto.Marshal(send)
	if err != nil {
		log.Println("Marshal SGUserAuth error:", err)
		return
	}
	oct := &common.Octets{}
	oct.MarshalUint32(uint32(len(data)))
	oct.MarshalUint32(2)
	oct.MarshalBytesOnly(data)
	msg.Global.Send(oct.GetBuf())
}
