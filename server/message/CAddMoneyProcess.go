package message

import (
	"gameproject/server/cacheMgr"
	"gameproject/server/lockMgr"
	"gameproject/server/protocol"
	"log"
	"strconv"
)

type CAddMoneyProcess struct {
	CAddMoney
}

func (this *CAddMoneyProcess) Process() {
	k := "MONEY_" + strconv.FormatUint(this.RoleId, 10)
	lockMgr.Lock(k)
	defer lockMgr.Unlock(k)

	v := cacheMgr.GetKV(k)
	sendInfo := &protocol.SMoneyInfo{}
	sendInfo.RoleId = this.RoleId
	if v == "" {
		sendInfo.Total = 0
	} else {
		i, _ := strconv.Atoi(v)
		sendInfo.Total = uint32(i)
	}

	sendInfo.Total += this.Num
	cacheMgr.SetKV(k, strconv.Itoa(int(sendInfo.Total)))
	err := this.Send(sendInfo)
	if err != nil {
		log.Panic("CAddMoneyProcess Send SMoneyInfo error:", err)
	}
}
