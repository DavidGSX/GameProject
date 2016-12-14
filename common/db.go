package common

import (
	"log"

	"github.com/mediocregopher/radix.v2/pool"
)

var db *pool.Pool

func GetDBPool() *pool.Pool {
	if db == nil {
		var err error
		db, err = pool.New("tcp", "127.0.0.1:6379", 10)
		if err != nil {
			log.Panic(err)
		}
	}
	return db
}
