package main

import (
	"gameproject/common"
	"gameproject/common/cache"
	"gameproject/world/config"
	"gameproject/world/manager"
	"gameproject/world/message"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	// 配置文件初始化
	cfg := config.GetConfig()
	cfg.Show()

	// 服务器协议初始化
	message.Init()

	// 数据缓存和数据库连接池初始化
	params := make(map[string]interface{}, 0)
	params["ip"] = cfg.DBIP
	params["port"] = cfg.DBPort
	params["dbName"] = "World"
	cache.CacheInit(cache.MongoDB, 1, params)
	<-time.After(2e9) // 2秒初始化时间

	// 锁管理器的初始化
	common.LockMgrInit()
	<-time.After(2e9) // 3秒初始化时间

	go manager.ServerMgrInit(cfg)
	go manager.JMXMgrInit(cfg, &wg)

	wg.Add(1)
	wg.Wait()
}
