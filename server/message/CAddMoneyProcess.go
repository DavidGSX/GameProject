package message

import (
	"log"
)

type CAddMoneyProcess struct {
	msg *CAddMoney
}

func (this *CAddMoneyProcess) Process(msg *CAddMoney) {
	this.msg = msg
	log.Println("to do CAddMoneyProcess")
}
