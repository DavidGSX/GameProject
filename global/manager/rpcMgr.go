package manager

import (
	"gameproject/global/config"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"strconv"
)

type UniqName int

func (this *UniqName) Exist(args *string, replay *bool) error {
	ret := DBGetKV("UniqName_" + *args)
	if ret == "true" {
		*replay = true
	} else {
		*replay = false
	}
	return nil
}

func (this *UniqName) Insert(args *string, replay *bool) error {
	DBSetKV("UniqName_"+*args, "true")
	*replay = true
	return nil
}

func (this *UniqName) Delete(args *string, replay *bool) error {
	DBDelKV("UniqName_" + *args)
	*replay = true
	return nil
}

func InitRPC(cfg *config.GlobalConfig) {
	ip := cfg.BaseConfig.RPCIP
	port := cfg.BaseConfig.RPCPort

	l, err := net.Listen("tcp", ip+":"+strconv.Itoa(int(port)))
	if err != nil {
		log.Fatal("RPC Listen Error:", err)
	}
	log.Println("RPC Listen ", ip, port)

	rpc.Register(new(UniqName))
	rpc.HandleHTTP()
	http.Serve(l, nil)
}
