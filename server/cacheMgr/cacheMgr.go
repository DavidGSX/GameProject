package cacheMgr

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

var dbCacheMap map[string]*ValueInfo
var dbCacheMapLock sync.Mutex

func CacheInit(cfg *config.ServerConfig) {
	dbCacheMapLock.Lock()
	defer dbCacheMapLock.Unlock()

	dbCacheMap = make(map[string]*ValueInfo)
	dbInit(cfg)

	go cleanCacheTicker()
}

func GetKV(k string) string {
	dbCacheMapLock.Lock()
	defer func() {
		dbCacheMapLock.Unlock()
		if err := recover(); err != nil {
			log.Println("Cache GetKV ", err, " Key:", k)
		}
	}()

	info, ok := dbCacheMap[k]
	if ok == false {
		info = new(ValueInfo)
		info.v = dbGetKV(k)
		dbCacheMap[k] = info
	}
	info.t = time.Now()
	return info.v
}

func SetKV(k, v string) {
	dbCacheMapLock.Lock()
	defer func() {
		dbCacheMapLock.Unlock()
		if err := recover(); err != nil {
			log.Println("Cache SetKV ", err, " Key:", k, " Value:", v)
		}
	}()
	dbSetKV(k, v)

	info, ok := dbCacheMap[k]
	if ok == false {
		info = new(ValueInfo)
	}
	info.v = v
	info.t = time.Now()
	dbCacheMap[k] = info
}

// 避免缓存过大，每10分钟清理下超过1小时的缓存，清理时间不能超过10ms
func cleanCacheTicker() {
	ticker := time.Tick(10 * time.Minute)
	for now := range ticker {
		cleanCache(now)
	}
}

func cleanCache(now time.Time) {
	dbCacheMapLock.Lock()
	defer dbCacheMapLock.Unlock()

	count := 0
	for k, v := range dbCacheMap {
		if time.Since(v.t) > 60*time.Minute {
			delete(dbCacheMap, k)
			count++
		}
		if time.Since(now) > 10*time.Millisecond {
			break
		}
	}
	log.Println("Clean Cache Use:", time.Since(now), " Size:", len(dbCacheMap), " Count:", count)
}
