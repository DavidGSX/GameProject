package message

import (
	"log"
)

type SAddMoneyProcess struct {
	msg *SAddMoney
}

func (this *SAddMoneyProcess) Process(msg *SAddMoney) {
	this.msg = msg
	log.Println("to do SAddMoneyProcess")
}
