package message

import (
	"gameproject/common"
	"log"
)

type S2WServerStartProcess struct {
	msg   *S2WServerStart
	trans *common.Trans
}

func (this *S2WServerStartProcess) Process() bool {
	defer func() {
		if err := recover(); err != nil {
			log.Println("S2WServerStartProcess Error:", err)
		}
	}()

	this.msg.Gets().SetZoneId(this.msg.ZoneId)

	sendInfo := &W2SServerStartRes{}
	err := this.msg.Send(sendInfo)
	if err != nil {
		log.Panic("S2WServerStartProcess Send W2SServerStartRes error:", err)
	}
	return true
}
