package main

import (
	"gameproject/server/cache"
	"gameproject/server/config"
	"gameproject/server/manager"
	"gameproject/server/message"
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
	cache.CacheInit(cfg)

	<-time.After(5e9) // 5秒初始化时间

	go manager.LinkMgrInit(cfg)
	go manager.GlobalMgrInit(cfg)
	go manager.JMXInit(cfg, &wg)

	wg.Add(1)
	wg.Wait()
}
