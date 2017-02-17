package csproc

import "gameproject/server/client/msgMgr"

func Init() {
	msgMgr.AddProc("CUserLogin", new(CUserLoginProcess))
	msgMgr.AddProc("CRoleList", new(CRoleListProcess))
	msgMgr.AddProc("CCreateRole", new(CCreateRoleProcess))
	msgMgr.AddProc("CEnterWorld", new(CEnterWorldProcess))
	msgMgr.AddProc("CAddMoney", new(CAddMoneyProcess))
	msgMgr.AddProc("CAddLevel", new(CAddLevelProcess))
	msgMgr.AddProc("CReqServerRoleInfos", new(CReqServerRoleInfosProcess))
}
