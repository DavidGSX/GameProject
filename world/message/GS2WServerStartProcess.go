package message

import (
	"gameproject/world/transMgr"
	"log"
)

type GS2WServerStartProcess struct {
	msg   *GS2WServerStart
	trans *transMgr.Trans
}

func (this *GS2WServerStartProcess) Process() bool {
	defer func() {
		if err := recover(); err != nil {
			log.Println("GS2WServerStartProcess Error:", err)
		}
	}()

	this.msg.Gets().SetZoneId(this.msg.ZoneId)

	sendInfo := &W2GSServerStartRes{}
	err := this.msg.Send(sendInfo)
	if err != nil {
		log.Panic("GS2WServerStartProcess Send W2GSServerStartRes error:", err)
	}
	return true
}
