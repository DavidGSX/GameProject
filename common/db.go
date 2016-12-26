package common

import (
	"log"
	"sync"

	"github.com/DavidGSX/gossdb"
)

var pool *gossdb.Connectors
var l sync.Mutex

func GetDBPool() *gossdb.Connectors {
	l.Lock()
	defer l.Unlock()
	if pool == nil {
		var err error
		pool, err = gossdb.NewPool(&gossdb.Config{
			Host:             "127.0.0.1",
			Port:             8888,
			MinPoolSize:      100,
			MaxPoolSize:      200,
			AcquireIncrement: 10,
		})
		if err != nil {
			log.Panic(err)
		}
		log.Println("create db pool!")
	}
	return pool
}

func SetKV(k, v string) {
	c, err := GetDBPool().NewClient()
	if err != nil {
		log.Panic(err)
	}
	defer c.Close()

	c.Set(k, v)
}
