package manager

import (
	"gameproject/global/config"
	"log"

	"github.com/DavidGSX/gossdb"
)

var dbPool *gossdb.Connectors

func DBInit(cfg *config.GlobalConfig) {
	ip := cfg.BaseConfig.DBIP
	port := cfg.BaseConfig.DBPort
	minPoolSize := cfg.BaseConfig.MinPoolSize
	maxPoolSize := cfg.BaseConfig.MaxPoolSize
	acqIncrement := cfg.BaseConfig.AcquireIncrement

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

func DBSetKV(k, v string) {
	c, err := dbPool.NewClient()
	if err != nil {
		log.Panic(err)
	}
	defer c.Close()

	c.Set(k, v)
}

func DBGetKV(k string) string {
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

func DBDelKV(k string) {
	c, err := dbPool.NewClient()
	if err != nil {
		log.Panic(err)
	}
	defer c.Close()

	c.Del(k)
}
