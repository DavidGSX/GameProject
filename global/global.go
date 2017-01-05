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
	<-time.After(5e9) // 5秒初始化时间

	go manager.InitHttpCallback(cfg)
	go manager.InitAuthor(cfg)
	go manager.InitJMX(cfg, &wg)

	wg.Add(1)
	wg.Wait()
}
