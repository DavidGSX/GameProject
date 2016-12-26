package main

import (
	"gameproject/global/config"
	"gameproject/global/manager"
	"sync"
)

var wg sync.WaitGroup

func main() {
	cfg := config.GetConfig()
	manager.GetPlatMgr().Init(cfg)

	go manager.InitHttpCallback(cfg)
	go manager.InitAuthor(cfg)
	go manager.InitJMX(cfg, &wg)

	wg.Add(1)
	wg.Wait()
}
