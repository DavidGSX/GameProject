package main

import (
	"gameproject/server/cacheMgr"
	"gameproject/server/config"
	"gameproject/server/jmxMgr"
	"gameproject/server/lockMgr"
	"gameproject/server/manager"
	"gameproject/server/message"
	"gameproject/server/rpcMgr"
	"log"
	"net/http"
	_ "net/http/pprof"
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

	// 服务器配置初始化
	cfg := config.GetConfig()
	cfg.Show()
	// 服务器协议初始化
	message.Init()
	// 数据缓存和数据库连接池初始化
	cacheMgr.CacheInit(cfg)
	<-time.After(2e9) // 2秒初始化时间
	// 锁管理器的初始化
	lockMgr.LockMgrInit()
	// RPC初始化
	rpcMgr.RPCInit(cfg)
	<-time.After(3e9) // 2秒初始化时间

	go manager.LinkMgrInit(cfg)
	go manager.GlobalMgrInit(cfg)
	go manager.WorldMgrInit(cfg)
	go jmxMgr.JMXInit(cfg, &wg)

	wg.Add(1)
	wg.Wait()
}
