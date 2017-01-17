package main

import (
	"gameproject/server/cacheMgr"
	"gameproject/server/config"
	"gameproject/server/jmxMgr"
	"gameproject/server/lockMgr"
	"gameproject/server/manager"
	"gameproject/server/message"
	"gameproject/server/rpcMgr"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	// 服务器配置初始化
	cfg := config.GetConfig()
	cfg.Show()
	// 服务器协议初始化
	message.Init()
	// 数据缓存和数据库连接池初始化
	cacheMgr.CacheInit(cfg)
	// 锁管理器的初始化
	lockMgr.LockMgrInit()
	// RPC初始化
	rpcMgr.RPCInit(cfg)

	<-time.After(3e9) // 3秒初始化时间

	go manager.LinkMgrInit(cfg)
	go manager.GlobalMgrInit(cfg)
	go jmxMgr.JMXInit(cfg, &wg)

	wg.Add(1)
	wg.Wait()
}
