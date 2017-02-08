package cacheMgr

import (
	"gameproject/world/config"
	"log"

	"github.com/DavidGSX/gossdb"
)

var ssdbPool *gossdb.Connectors

func ssdbInit(cfg *config.WorldConfig) {
	ip := cfg.DBIP
	port := cfg.DBPort
	minPoolSize := cfg.MinPoolSize
	maxPoolSize := cfg.MaxPoolSize
	acqIncrement := cfg.AcqIncrement

	var err error
	ssdbPool, err = gossdb.NewPool(&gossdb.Config{
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

func ssdbSetKV(k, v string) {
	c, err := ssdbPool.NewClient()
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

func ssdbGetKV(k string) string {
	c, err := ssdbPool.NewClient()
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
