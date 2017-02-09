package csmsg

import (
	"gameproject/common"

	"github.com/golang/protobuf/proto"
)

// 避免与协议的函数名称重复，函数的命名有点特殊
type MsgInfo interface {
	Clone() MsgInfo
	MsgType() uint32
	GetMsg() proto.Message
	Setr(r uint64)
	Getr() uint64
	Setl(s ISend)
	Getl() ISend
	Setg(s ISend)
	Getg() ISend
	Unmarshal(data []byte) error
	Send(MsgInfo) error
	Process(t *common.Trans) bool
}

type ISend interface {
	Send(x []byte)
	SetUserId(u string)
	GetUserId() string
}

var MsgInfos map[int]MsgInfo

func Init() {
	MsgInfos = make(map[int]MsgInfo)
	MsgInfos[1001] = new(CUserLogin)
	MsgInfos[1002] = new(SUserLogin)
	MsgInfos[1003] = new(CRoleList)
	MsgInfos[1004] = new(SRoleList)
	MsgInfos[1005] = new(CCreateRole)
	MsgInfos[1006] = new(SCreateRole)
	MsgInfos[1007] = new(CAddMoney)
	MsgInfos[1008] = new(SMoneyInfo)
}

func GetMsg(t int) MsgInfo {
	if MsgInfos == nil {
		return nil
	}
	return MsgInfos[t]
}
