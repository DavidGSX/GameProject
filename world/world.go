package main

import (
	"gameproject/common"
	"gameproject/world/config"
	"gameproject/world/db/cacheMgr"
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
	// 数据缓存和数据库连接池初始化
	cacheMgr.CacheInit(cfg)
	<-time.After(2e9) // 2秒初始化时间
	// 锁管理器的初始化
	common.LockMgrInit()
	<-time.After(2e9) // 3秒初始化时间

	go manager.ServerMgrInit(cfg)
	go manager.JMXMgrInit(cfg, &wg)

	wg.Add(1)
	wg.Wait()
}
