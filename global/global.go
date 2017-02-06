package main

import (
	"gameproject/global/config"
	"gameproject/global/manager"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	cfg := config.GetConfig()
	cfg.Show()
	manager.GetPlatMgr().Init(cfg)
	//manager.SSDBInit(cfg)
	manager.MongoDBInit()
	<-time.After(3e9) // 3秒初始化时间

	go manager.CallbackMgrInit(cfg)
	go manager.ServerMgrInit(cfg)
	go manager.RPCMgrInit(cfg)
	go manager.JMXMgrInit(cfg, &wg)

	wg.Add(1)
	wg.Wait()
}
