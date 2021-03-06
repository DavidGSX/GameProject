package csproc

import (
	"gameproject/common"
	"gameproject/global/protocol"
	"gameproject/server/client/csmsg"
	"gameproject/server/client/msgMgr"
	"log"

	"github.com/golang/protobuf/proto"
)

type CUserLoginProcess struct {
	msg   *csmsg.CUserLogin
	trans *common.Trans
}

func (this *CUserLoginProcess) Clone() msgMgr.IProcess {
	return new(CUserLoginProcess)
}

func (this *CUserLoginProcess) SetMsg(m msgMgr.MsgInfo) {
	this.msg = m.(*csmsg.CUserLogin)
}

func (this *CUserLoginProcess) SetTrans(t *common.Trans) {
	this.trans = t
}

func (this *CUserLoginProcess) Process() bool {
	defer func() {
		if err := recover(); err != nil {
			log.Println("CUserLoginProcess Error:", err)
		}
	}()

	userId := this.msg.UserId
	token := this.msg.Token

	//log.Println("UserId:", userId, " Token:", token)
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
