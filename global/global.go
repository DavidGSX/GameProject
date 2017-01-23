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

	go manager.InitHttpCallback(cfg)
	go manager.InitAuthor(cfg)
	go manager.InitRPC(cfg)
	go manager.InitJMX(cfg, &wg)

	wg.Add(1)
	wg.Wait()
}
