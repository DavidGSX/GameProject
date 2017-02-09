package message

import (
	"gameproject/common"

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
	Process(t *common.Trans) bool
}

type ISend interface {
	Send(b []byte)
	SendByZoneIds(zoneIds []uint32, b []byte)
	SetZoneId(z uint32)
	GetZoneId() uint32
}

func GetMsgByte(msg MsgInfo) (error, []byte) {
	data, err := proto.Marshal(msg.GetMsg())
	if err != nil {
		return err, nil
	}
	oct := &common.Octets{}
	oct.MarshalUint32(uint32(len(data)))
	oct.MarshalUint32(msg.MsgType())
	oct.MarshalBytesOnly(data)
	return nil, oct.GetBuf()
}

var MsgInfos map[int]MsgInfo

func Init() {
	MsgInfos = make(map[int]MsgInfo)
	MsgInfos[101] = new(S2WServerStart)
	MsgInfos[102] = new(W2SServerStartRes)
	MsgInfos[103] = new(S2WSendInfo)
	MsgInfos[104] = new(W2SSendInfo)
	MsgInfos[105] = new(W2SDispatch)
	MsgInfos[106] = new(W2SBroadcast)
	MsgInfos[107] = new(S2WRoleInfoChange)
	MsgInfos[108] = new(S2WRoleInfoRequest)
	MsgInfos[109] = new(W2SRoleInfoResponse)
}

func GetMsg(t int) MsgInfo {
	if MsgInfos == nil {
		return nil
	}
	return MsgInfos[t]
}
