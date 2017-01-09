package message

import (
	"gameproject/server/cacheMgr"
	"gameproject/server/protocol"
	"log"
)

type CRoleListProcess struct {
	CRoleList
}

func (this *CRoleListProcess) Process() {
	k := "USER" + this.Getl().GetUserId()
	v := cacheMgr.GetKV(k)

	sendInfo := &protocol.SRoleList{}
	if v != "" {
		// Decode DB Data
	}
	sendInfo.PreLoginRoleId = 1
	err := this.Send(sendInfo)
	if err != nil {
		log.Panic("CRoleListProcess Send SRoleList error:", err)
	}
}
