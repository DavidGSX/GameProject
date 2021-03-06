package common

import (
	"log"
	"sync"
	"time"
)

type LockInfo struct {
	l sync.Mutex // 锁
	t time.Time  // 加锁的时间
	b bool       // 是否加锁
	s string     // 锁的名称
}

func (this *LockInfo) Lock(name string) {
	this.l.Lock()
	this.t = time.Now()
	this.b = true
	this.s = name
}

func (this *LockInfo) Unlock() {
	if this.b {
		this.b = false
		this.l.Unlock()
	} else {
		log.Println("Unlock Error, " + this.s + " is not Locked!")
	}
}

var dbLockMap map[string]*LockInfo
var dbLockMapLock sync.Mutex

func LockMgrInit() {
	dbLockMapLock.Lock()
	defer dbLockMapLock.Unlock()

	dbLockMap = make(map[string]*LockInfo)
	go cleanLockTicker()
}

func getLock(name string) *LockInfo {
	dbLockMapLock.Lock()
	defer dbLockMapLock.Unlock()

	if name == "" {
		log.Panic("LockMgr.Lock name is nil")
	}
	dbLock, ok := dbLockMap[name]
	if ok {
		return dbLock
	} else {
		dbLockMap[name] = new(LockInfo)
		return dbLockMap[name]
	}
}

func Lock(names ...string) {
	names = SortAndRemoveEmptyDuplicates(names)
	for _, v := range names {
		getLock(v).Lock(v)
	}
	//log.Println("lock ", names)
}

func Unlock(names ...string) {
	names = SortAndRemoveEmptyDuplicates(names)
	for _, v := range names {
		getLock(v).Unlock()
	}
	//log.Print("unlock", names)
}

// 为了打断死锁，每5秒钟清理下超过5秒钟的锁和超过10分钟的缓存
func cleanLockTicker() {
	ticker := time.Tick(5 * time.Second)
	for now := range ticker {
		cleanLock(now)
	}
}

func cleanLock(now time.Time) {
	dbLockMapLock.Lock()
	defer dbLockMapLock.Unlock()
	count := 0
	for k, v := range dbLockMap {
		if v.b == true && time.Since(v.t) > 5*time.Second {
			v.Unlock()
			log.Println("Clean DeadLock ", v.s)
		}
		// 清理冗余的数据，避免遍历map时间过长
		if v.b == false && time.Since(v.t) > 10*time.Minute {
			delete(dbLockMap, k)
			count++
		}
	}
	log.Println("Clean Lock Use:", time.Since(now), " Size:", len(dbLockMap), " Count:", count)
}
