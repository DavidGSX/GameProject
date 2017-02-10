package main

import (
	"gameproject/common"
	"gameproject/common/cache"
	"gameproject/server/client/csmsg"
	"gameproject/server/config"
	"gameproject/server/manager"
	"gameproject/server/rpcMgr"
	"gameproject/server/world/swmsg"
	"log"
	"net/http"
	_ "net/http/pprof"
	"strconv"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	// 服务器性能分析
	// 分析CPU占用 go tool pprof http://localhost:6060/debug/pprof/profile
	// 分析内存占用 go tool pprof http://localhost:6060/debug/pprof/heap
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	// 配置文件初始化
	cfg := config.GetConfig()
	cfg.Show()

	// 服务器协议初始化
	csmsg.Init()
	swmsg.Init()

	// 数据缓存和数据库连接池初始化
	params := make(map[string]interface{}, 0)
	params["ip"] = cfg.DBConfig.DBIP
	params["port"] = cfg.DBConfig.DBPort
	params["dbName"] = "Game_" + strconv.Itoa(int(cfg.BaseConfig.ZoneId))
	cache.CacheInit(cache.MongoDB, cfg.BaseConfig.ZoneId, params)
	<-time.After(2e9) // 2秒初始化时间

	// 锁管理器的初始化
	common.LockMgrInit()

	// RPC初始化
	rpcMgr.RPCInit(cfg)
	<-time.After(3e9) // 2秒初始化时间

	go manager.LinkMgrInit(cfg)
	go manager.GlobalMgrInit(cfg)
	go manager.WorldMgrInit(cfg)
	go manager.JMXInit(cfg, &wg)

	wg.Add(1)
	wg.Wait()
}
