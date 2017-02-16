package msgMgr

import (
	"gameproject/common"
	"log"

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

func AddMsg(t int, msg MsgInfo) {
	if MsgInfos == nil {
		MsgInfos = make(map[int]MsgInfo)
	}
	if _, ok := MsgInfos[t]; ok {
		log.Panic("Duplicate Msg Type", t)
	}
	MsgInfos[t] = msg
}

func GetMsg(t int) MsgInfo {
	if MsgInfos == nil {
		return nil
	}
	return MsgInfos[t].Clone()
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

type IProcess interface {
	Clone() IProcess
	SetMsg(m MsgInfo)
	SetTrans(t *common.Trans)
	Process() bool
}

var ProcInfos map[string]IProcess

func AddProc(s string, p IProcess) {
	if ProcInfos == nil {
		ProcInfos = make(map[string]IProcess)
	}
	if _, ok := ProcInfos[s]; ok {
		log.Panic("Duplicate Proc Type", s)
	}
	ProcInfos[s] = p
}

func GetProc(s string) IProcess {
	if ProcInfos == nil {
		return nil
	}
	return ProcInfos[s].Clone()
}

