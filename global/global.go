package main

import (
	"gameproject/common/cache"
	"gameproject/global/config"
	"gameproject/global/manager"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	// 配置文件初始化
	cfg := config.GetConfig()
	cfg.Show()

	// 渠道管理初始化
	manager.GetPlatMgr().Init(cfg)

	// 数据缓存和数据库连接池初始化
	params := make(map[string]interface{}, 0)
	params["ip"] = cfg.BaseConfig.DBIP
	params["port"] = cfg.BaseConfig.DBPort
	params["dbName"] = "Global"
	cache.CacheInit(cache.MongoDB, 1, params)
	<-time.After(2e9) // 2秒初始化时间

	go manager.CallbackMgrInit(cfg)
	go manager.ServerMgrInit(cfg)
	go manager.RPCMgrInit(cfg)
	go manager.JMXMgrInit(cfg, &wg)

	wg.Add(1)
	wg.Wait()
}
