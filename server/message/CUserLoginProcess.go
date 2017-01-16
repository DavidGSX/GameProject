package message

import (
	"gameproject/common"
	"gameproject/global/protocol"
	"gameproject/server/transMgr"
	"log"

	"github.com/golang/protobuf/proto"
)

type CUserLoginProcess struct {
	msg   *CUserLogin
	trans *transMgr.Trans
}

func (this *CUserLoginProcess) Process() bool {
	defer func() {
		if err := recover(); err != nil {
			log.Println("CUserLoginProcess Error:", err)
		}
	}()

	userId := this.msg.UserId
	token := this.msg.Token

	log.Println("UserId:", userId, " Token:", token)
	this.msg.Getl().SetUserId(userId)

	send := &protocol.SGUserAuth{}
	send.UserId = userId
	send.Token = token
	data, err := proto.Marshal(send)
	if err != nil {
		log.Println("CUserLoginProcess Marshal SGUserAuth error:", err)
		return false
	}
	oct := &common.Octets{}
	oct.MarshalUint32(uint32(len(data)))
	oct.MarshalUint32(2)
	oct.MarshalBytesOnly(data)
	this.msg.Getg().Send(oct.GetBuf())
	return true
}
