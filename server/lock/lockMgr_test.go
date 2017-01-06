package lock

import (
	"log"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func TestLock(t *testing.T) {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	c := make(chan int, 1)
	c <- 0
	for i := 0; i < 1e6; i++ {
		strs := []string{""}
		for i := 0; i < 3; i++ {
			strs = append(strs, strconv.Itoa(r.Intn(10)))
		}
		go goLock(c, strs...)
	}

	<-time.After(5e9) // 5秒后退出
	log.Println(<-c)
}

func goLock(c chan int, strs ...string) {
	GetLockMgr().Lock(strs...)
	GetLockMgr().Unlock(strs...)
	c <- 1 + <-c
}
