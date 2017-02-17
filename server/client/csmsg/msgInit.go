package csmsg

import "gameproject/server/client/msgMgr"

func Init() {
	msgMgr.AddMsg(1001, new(CUserLogin))
	msgMgr.AddMsg(1002, new(SUserLogin))
	msgMgr.AddMsg(1003, new(CRoleList))
	msgMgr.AddMsg(1004, new(SRoleList))
	msgMgr.AddMsg(1005, new(CCreateRole))
	msgMgr.AddMsg(1006, new(SCreateRole))
	msgMgr.AddMsg(1007, new(CEnterWorld))
	msgMgr.AddMsg(1008, new(SEnterWorld))
	msgMgr.AddMsg(1009, new(CAddMoney))
	msgMgr.AddMsg(1010, new(SMoneyInfo))
	msgMgr.AddMsg(1011, new(CAddLevel))
	msgMgr.AddMsg(1012, new(SLevelInfo))
	msgMgr.AddMsg(1013, new(CReqServerRoleInfos))
	msgMgr.AddMsg(1014, new(SServerRoleInfos))
}
