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
	Setw(w ISend)
	Getw() ISend
	Unmarshal(data []byte) error
	Send2Link(MsgInfo) error
	Process(t *common.Trans) bool
}

type ISend interface {
	Send(x []byte)
	SetUserId(u string)
	GetUserId() string
	SetRoleId(r uint64)
	GetRoleId() uint64
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
	MsgInfos[1007] = new(CEnterWorld)
	MsgInfos[1008] = new(SEnterWorld)
	MsgInfos[1009] = new(CAddMoney)
	MsgInfos[1010] = new(SMoneyInfo)
	MsgInfos[1011] = new(CAddLevel)
	MsgInfos[1012] = new(SLevelInfo)
	MsgInfos[1013] = new(CReqServerRoleInfos)
	MsgInfos[1014] = new(SServerRoleInfos)
}

func GetMsg(t int) MsgInfo {
	if MsgInfos == nil {
		return nil
	}
	return MsgInfos[t]
}

func MarshalMsg(msg MsgInfo) ([]byte, error) {
	data, err := proto.Marshal(msg.GetMsg())
	if err != nil {
		return nil, err
	}
	oct := &common.Octets{}
	oct.MarshalUint32(uint32(len(data)))
	oct.MarshalUint32(msg.MsgType())
	oct.MarshalBytesOnly(data)
	return oct.GetBuf(), nil
}
