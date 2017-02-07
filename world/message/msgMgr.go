package message

import (
	"gameproject/world/transMgr"

	"github.com/golang/protobuf/proto"
)

// 避免与协议的函数名称重复，函数的命名有点特殊
type MsgInfo interface {
	Clone() MsgInfo
	MsgType() uint32
	GetMsg() proto.Message
	Sets(s ISend)
	Gets() ISend
	Unmarshal(data []byte) error
	Send(MsgInfo) error
	Process(t *transMgr.Trans) bool
}

type ISend interface {
	Send(x []byte)
	SetZoneId(z uint32)
	GetZoneId() uint32
}

var MsgInfos map[int]MsgInfo

func Init() {
	MsgInfos = make(map[int]MsgInfo)
	MsgInfos[101] = new(GS2WServerStart)
	MsgInfos[102] = new(W2GSServerStartRes)
	MsgInfos[103] = new(GS2WSendInfo)
	MsgInfos[104] = new(W2GSSendInfo)
	MsgInfos[105] = new(W2GSDispatch)
	MsgInfos[106] = new(W2GSBroadcast)
}

func GetMsg(t int) MsgInfo {
	if MsgInfos == nil {
		return nil
	}
	return MsgInfos[t]
}
