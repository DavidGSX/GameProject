package cacheMgr

import (
	"gameproject/server/config"
	"log"

	"github.com/DavidGSX/gossdb"
)

var dbPool *gossdb.Connectors

func dbInit(cfg *config.ServerConfig) {
	ip := cfg.DBConfig.DBIP
	port := cfg.DBConfig.DBPort
	minPoolSize := cfg.DBConfig.MinPoolSize
	maxPoolSize := cfg.DBConfig.MaxPoolSize
	acqIncrement := cfg.DBConfig.AcquireIncrement

	var err error
	dbPool, err = gossdb.NewPool(&gossdb.Config{
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

func dbSetKV(k, v string) {

	c, err := dbPool.NewClient()
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

func dbGetKV(k string) string {
	c, err := dbPool.NewClient()
	if err != nil {
		log.Panic(err)
	}
	defer c.Close()

	v, err := c.Get(k)
	if err != nil {
		log.Panic(err)
	}
	return v.String()
}
