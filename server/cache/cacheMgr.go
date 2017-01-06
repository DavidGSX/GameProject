package cache

import (
	"gameproject/server/config"
	"log"
	"sync"
	"time"
)

type ValueInfo struct {
	v string
	t time.Time
}

var cache map[string]*ValueInfo
var cacheLock sync.Mutex

func CacheInit(cfg *config.ServerConfig) {
	cacheLock.Lock()
	defer cacheLock.Unlock()

	cache = make(map[string]*ValueInfo)
	dbInit(cfg)

	go cleanCacheTicker()
}

func GetKV(k string) string {
	cacheLock.Lock()
	defer func() {
		cacheLock.Unlock()
		if err := recover(); err != nil {
			log.Println("Cache GetKV ", err, " Key:", k)
		}
	}()

	info, ok := cache[k]
	if ok == false {
		info = new(ValueInfo)
		info.v = dbGetKV(k)
		cache[k] = info
	}
	info.t = time.Now()
	return info.v
}

func SetKV(k, v string) {
	cacheLock.Lock()
	defer func() {
		cacheLock.Unlock()
		if err := recover(); err != nil {
			log.Println("Cache SetKV ", err, " Key:", k, " Value:", v)
		}
	}()
	dbSetKV(k, v)

	info, ok := cache[k]
	if ok == false {
		info = new(ValueInfo)
	}
	info.v = v
	info.t = time.Now()
	cache[k] = info
}

// 避免缓存过大，每10分钟清理下超过1小时的缓存，清理时间不能超过10ms
func cleanCacheTicker() {
	ticker := time.Tick(10 * time.Minute)
	for now := range ticker {
		cleanCache(now)
	}
}

func cleanCache(now time.Time) {
	cacheLock.Lock()
	defer cacheLock.Unlock()

	count := 0
	for k, v := range cache {
		if time.Since(v.t) > 60*time.Minute {
			delete(cache, k)
			count++
		}
		if time.Since(now) > 10*time.Millisecond {
			break
		}
	}
	log.Println(" Clean Count:", count, " Use:", time.Since(now))
}
