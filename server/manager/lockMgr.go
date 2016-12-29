package manager

import (
	"log"
	"sync"
)

type LockMgr struct {
	dbLockMap map[string]*sync.Mutex
	mapLock   sync.Mutex
}

var lockMgr *LockMgr
var lockMgrLock sync.Mutex

func GetLockMgr() *LockMgr {
	lockMgrLock.Lock()
	defer lockMgrLock.Unlock()

	if lockMgr == nil {
		lockMgr = &LockMgr{}
		lockMgr.dbLockMap = make(map[string]*sync.Mutex)
	}
	return lockMgr
}

func (this *LockMgr) GetLock(name string) *sync.Mutex {
	this.mapLock.Lock()
	defer this.mapLock.Unlock()

	if name == "" {
		log.Panic("LockMgr.Lock name is nil")
	}
	dbLock, ok := this.dbLockMap[name]
	if ok {
		return dbLock
	} else {
		this.dbLockMap[name] = new(sync.Mutex)
		return this.dbLockMap[name]
	}
}

func (this *LockMgr) Lock(name string) {
	this.GetLock(name).Lock()
}

func (this *LockMgr) Unlock(name string) {
	this.mapLock.Lock()
	defer this.mapLock.Unlock()

	if name == "" {
		log.Panic("LockMgr.Unlock name is nil")
	}
	dbLock, ok := this.dbLockMap[name]
	if ok {
		dbLock.Unlock()
	} else {
		log.Panic("LockMgr.Unlock not find name:", name)
	}
}
