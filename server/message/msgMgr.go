package message

type MsgInfo interface {
	Clone() MsgInfo
	SetRoleId(r uint64)
	SetLink(s ISend)
	SetGlobal(s ISend)
	Unmarshal(data []byte) error
	Process()
}

type ISend interface {
	Send(x []byte)
}

var MsgInfos map[int]MsgInfo

func Init() {
	MsgInfos = make(map[int]MsgInfo)
	MsgInfos[1001] = new(CUserLogin)
	MsgInfos[1002] = new(SUserLogin)
	MsgInfos[1003] = new(CRoleList)
	MsgInfos[1004] = new(SRoleList)
	MsgInfos[1005] = new(CAddMoney)
	MsgInfos[1006] = new(SAddMoney)
}

func GetMsg(t int) MsgInfo {
	if MsgInfos == nil {
		return nil
	}
	return MsgInfos[t]
}

