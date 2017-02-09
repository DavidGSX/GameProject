package cache

import (
	"log"
	"sync"
	"time"
)

type ValueInfo struct {
	v string
	t time.Time
}

type DBType int32

const (
	MongoDB DBType = 1
	SSDB    DBType = 2
)

var (
	dbCacheMap     map[string]*ValueInfo
	dbCacheMapLock sync.Mutex
	dbType         DBType
)

func CacheInit(dt DBType, zoneId uint32, params map[string]interface{}) {
	dbCacheMapLock.Lock()
	defer dbCacheMapLock.Unlock()

	dbCacheMap = make(map[string]*ValueInfo)
	SetZoneId(zoneId)

	dbType = dt
	if dbType == MongoDB {
		ip, ok1 := params["ip"]
		port, ok2 := params["port"]
		dbName, ok3 := params["dbName"]
		if ok1 == false || ok2 == false || ok3 == false {
			log.Panic("MongoDb Need ip, port, dbName")
		}
		mongoDBInit(ip.(string), port.(uint32), dbName.(string))
	} else if dbType == SSDB {
		ip, ok1 := params["ip"]
		port, ok2 := params["port"]
		minPoolSize, ok3 := params["minPoolSize"]
		maxPoolSize, ok4 := params["maxPoolSize"]
		acqIncrement, ok5 := params["acqIncrement"]
		if ok1 == false || ok2 == false || ok3 == false || ok4 == false || ok5 == false {
			log.Panic("SSDB Need ip, port, minPoolSize, maxPoolSize, acqIncrement")
		}
		ssdbInit(ip.(string), port.(uint32), minPoolSize.(uint32), maxPoolSize.(uint32), acqIncrement.(uint32))
	} else {
		log.Panic("Invalid DB Type:", dbType)
	}

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

		if dbType == MongoDB {
			info.v = mongoDBGetKV(k)
		} else if dbType == SSDB {
			info.v = ssdbGetKV(k)
		} else {
			log.Panic("Invalid DB Type:", dbType)
		}

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

	if dbType == MongoDB {
		mongoDBUpsertKV(k, v)
	} else if dbType == SSDB {
		ssdbSetKV(k, v)
	} else {
		log.Panic("Invalid DB Type:", dbType)
	}

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
