package message

import (
	"gameproject/common"
	"gameproject/server/db"
	"gameproject/server/lock"
	"gameproject/server/protocol"
	"log"
	"strconv"

	"github.com/golang/protobuf/proto"
)

type CAddMoneyProcess struct {
	msg *CAddMoney
}

func (this *CAddMoneyProcess) Process(msg *CAddMoney) {
	this.msg = msg

	k := "MONEY" + strconv.FormatUint(msg.Proto.RoleId, 10)
	lock.GetLockMgr().Lock(k)
	defer lock.GetLockMgr().Unlock(k)

	v := db.GetKV(k)
	sendInfo := &protocol.SMoneyInfo{}
	sendInfo.RoleId = msg.Proto.RoleId
	if v == "" {
		sendInfo.Total = 0
	} else {
		i, _ := strconv.Atoi(v)
		sendInfo.Total = uint32(i)
	}

	sendInfo.Total += msg.Proto.Num
	db.SetKV(k, strconv.Itoa(int(sendInfo.Total)))

	data, err := proto.Marshal(sendInfo)
	if err != nil {
		log.Panic("marshal error:", err)
	}
	oct := &common.Octets{}
	oct.MarshalUint32(uint32(len(data)))
	oct.MarshalUint32(1008)
	oct.MarshalBytesOnly(data)
	msg.Link.Send(oct.GetBuf())
}
