package rpcMgr

import (
	"gameproject/server/config"
	"log"
	"net/rpc"
	"strconv"
)

var rpcInfo *RpcInfo

type RpcInfo struct {
	coninfo string // ip:port
}

func RPCInit(cfg *config.ServerConfig) {
	ip := cfg.RPCConfig.RPCIP
	port := cfg.RPCConfig.RPCPort

	rpcInfo = new(RpcInfo)
	rpcInfo.coninfo = ip + ":" + strconv.Itoa(int(port))
}

func NameExist(name string) (ret bool) {
	client, err := rpc.DialHTTP("tcp", rpcInfo.coninfo)
	if err != nil {
		log.Panic("NameExist DialHTTP", err)
	}
	defer client.Close()

	err = client.Call("UniqName.Exist", &name, &ret)
	if err != nil {
		log.Panic("NameExist Call", err)
	}
	return ret
}

func NameInsert(name string) {
	client, err := rpc.DialHTTP("tcp", rpcInfo.coninfo)
	if err != nil {
		log.Panic("NameInsert DialHTTP", err)
	}
	defer client.Close()

	var ret bool
	err = client.Call("UniqName.Insert", &name, &ret)
	if err != nil {
		log.Panic("NameInsert Call", err)
	}
}

func NameDelete(name string) {
	client, err := rpc.DialHTTP("tcp", rpcInfo.coninfo)
	if err != nil {
		log.Panic("NameDelete DialHTTP", err)
	}
	defer client.Close()

	var ret bool
	err = client.Call("UniqName.Delete", &name, &ret)
	if err != nil {
		log.Panic("NameDelete Call", err)
	}
}
