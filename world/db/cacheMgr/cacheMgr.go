package cacheMgr

import (
	"gameproject/world/config"
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

func CacheInit(cfg *config.WorldConfig) {
	dbCacheMapLock.Lock()
	defer dbCacheMapLock.Unlock()

	dbCacheMap = make(map[string]*ValueInfo)
	//ssdbInit(cfg)
	mongoDBInit()

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
		//info.v = ssdbGetKV(k)
		info.v = mongoDBGetKV(k)
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
	//ssdbSetKV(k, v)
	mongoDBUpsertKV(k, v)

	info, ok := dbCacheMap[k]
	if ok == false {
		info = new(ValueInfo)
	}
	info.v = v
	info.t = time.Now()
	dbCacheMap[k] = info
}

// 为了避免缓存过大，每10秒清理下超过10分钟的缓存，清理时间不能超过30ms
func cleanCacheTicker() {
	ticker := time.Tick(10 * time.Second)
	for now := range ticker {
		cleanCache(now)
	}
}

func cleanCache(now time.Time) {
	dbCacheMapLock.Lock()
	defer dbCacheMapLock.Unlock()

	count := 0
	for k, v := range dbCacheMap {
		if time.Since(v.t) > 10*time.Minute {
			delete(dbCacheMap, k)
			count++
		}
		if time.Since(now) > 30*time.Millisecond {
			break
		}
	}
	log.Println("Clean Cache Use:", time.Since(now), " Size:", len(dbCacheMap), " Count:", count)
}
