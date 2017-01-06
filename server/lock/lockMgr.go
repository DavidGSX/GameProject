package lock

import (
	"gameproject/common"
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

func (this *LockMgr) getLock(name string) *sync.Mutex {
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

func (this *LockMgr) Lock(names ...string) {
	names = common.SortAndRemoveEmptyDuplicates(names)
	for _, v := range names {
		this.getLock(v).Lock()
	}
	//log.Println("lock ", names)
}

func (this *LockMgr) Unlock(names ...string) {
	names = common.SortAndRemoveEmptyDuplicates(names)
	for _, v := range names {
		this.getLock(v).Unlock()
	}
	//log.Print("unlock", names)
}
