package message

import (
	"gameproject/common"
	"log"
)

type S2WSendInfoProcess struct {
	msg   *S2WSendInfo
	trans *common.Trans
}

func (this *S2WSendInfoProcess) Process() bool {
	defer func() {
		if err := recover(); err != nil {
			log.Println("S2WSendInfoProcess Error:", err)
		}
	}()

	sendInfo := &W2SSendInfo{}
	sendInfo.ZoneId = this.msg.ZoneId
	sendInfo.UserId = this.msg.UserId
	sendInfo.Type = this.msg.Type
	sendInfo.Info = this.msg.Info

	err, data := GetMsgByte(sendInfo)
	if err != nil {
		log.Panic("S2WSendInfoProcess GetMsgByte W2SSendInfo error:", err)
	}
	this.msg.Gets().SendByZoneIds([]uint32{sendInfo.ZoneId}, data)
	return true
}
