package db

import (
	"gameproject/server/config"
	"log"
	"sync"

	"github.com/DavidGSX/gossdb"
)

var dbMgr *gossdb.Connectors
var dbMgrLock sync.Mutex

func GetDBMgr() *gossdb.Connectors {
	dbMgrLock.Lock()
	defer dbMgrLock.Unlock()

	return dbMgr
}

func DBMgrInit(cfg *config.ServerConfig) {

	ip := cfg.DBConfig.DBIP
	port := cfg.DBConfig.DBPort
	minPoolSize := cfg.DBConfig.MinPoolSize
	maxPoolSize := cfg.DBConfig.MaxPoolSize
	acqIncrement := cfg.DBConfig.AcquireIncrement

	dbMgrLock.Lock()
	defer dbMgrLock.Unlock()

	if dbMgr == nil {
		var err error
		dbMgr, err = gossdb.NewPool(&gossdb.Config{
			Host:             ip,
			Port:             int(port),
			MinPoolSize:      int(minPoolSize),
			MaxPoolSize:      int(maxPoolSize),
			AcquireIncrement: int(acqIncrement),
		})
		if err != nil {
			log.Panic(err)
		}
		log.Println("Create DB Pool [", minPoolSize, maxPoolSize, "] Success! ")
	}
}

func SetKV(k, v string) {
	c, err := GetDBMgr().NewClient()
	if err != nil {
		log.Panic(err)
	}
	defer func() {
		c.Close()
		if err := recover(); err != nil {
			log.Println("DB SetKV ", err, " Key:", k, " Value:", v)
		}
	}()

	c.Set(k, v)
}

func GetKV(k string) string {
	c, err := GetDBMgr().NewClient()
	if err != nil {
		log.Panic(err)
	}
	defer func() {
		c.Close()
		if err := recover(); err != nil {
			log.Println("DB GetKV ", err, " Key:", k)
		}
	}()

	v, err := c.Get(k)
	if err != nil {
		log.Panic(err)
	}
	return v.String()
}
