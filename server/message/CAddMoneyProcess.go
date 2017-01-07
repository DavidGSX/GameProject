package message

import (
	"gameproject/common"
	"gameproject/server/cacheMgr"
	"gameproject/server/lockMgr"
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
	lockMgr.Lock(k)
	defer lockMgr.Unlock(k)

	v := cacheMgr.GetKV(k)
	sendInfo := &protocol.SMoneyInfo{}
	sendInfo.RoleId = msg.Proto.RoleId
	if v == "" {
		sendInfo.Total = 0
	} else {
		i, _ := strconv.Atoi(v)
		sendInfo.Total = uint32(i)
	}

	sendInfo.Total += msg.Proto.Num
	cacheMgr.SetKV(k, strconv.Itoa(int(sendInfo.Total)))

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
