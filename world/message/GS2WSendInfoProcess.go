package message

import (
	"gameproject/world/transMgr"
	"log"
)

type GS2WSendInfoProcess struct {
	msg   *GS2WSendInfo
	trans *transMgr.Trans
}

func (this *GS2WSendInfoProcess) Process() bool {
	defer func() {
		if err := recover(); err != nil {
			log.Println("GS2WSendInfoProcess Error:", err)
		}
	}()

	sendInfo := &W2GSSendInfo{}
	sendInfo.ZoneId = this.msg.ZoneId
	sendInfo.UserId = this.msg.UserId
	sendInfo.Type = this.msg.Type
	sendInfo.Info = this.msg.Info

	zoneIds := make([]uint32, 0)
	zoneIds = append(zoneIds, sendInfo.ZoneId)
	err, data := GetMsgByte(sendInfo)
	if err != nil {
		log.Panic("GS2WSendInfoProcess GetMsgByte W2GSSendInfo error:", err)
	}
	this.msg.Gets().SendByZoneIds(zoneIds, data)
	return true
}
