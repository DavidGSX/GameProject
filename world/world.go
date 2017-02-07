package main

import (
	"gameproject/world/config"
	"gameproject/world/manager"
	"gameproject/world/message"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	cfg := config.GetConfig()
	cfg.Show()
	//manager.SSDBInit(cfg)
	manager.MongoDBInit()

	// 服务器协议初始化
	message.Init()

	<-time.After(3e9) // 3秒初始化时间

	go manager.ServerMgrInit(cfg)
	go manager.JMXMgrInit(cfg, &wg)

	wg.Add(1)
	wg.Wait()
}
