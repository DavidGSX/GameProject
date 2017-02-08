package lockMgr

import (
	"log"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func TestLock(t *testing.T) {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	LockMgrInit()
	c := make(chan int, 1)
	c <- 0
	for i := 0; i < 1e5; i++ {
		strs := []string{""}
		for i := 0; i < 10; i++ {
			strs = append(strs, strconv.Itoa(r.Intn(10000)))
		}
		go goLock(c, strs...)
	}

	<-time.After(1e10) // 10秒后退出
	log.Println(<-c)
}

func goLock(c chan int, strs ...string) {
	Lock(strs...)
	defer Unlock(strs...)
	<-time.After(1e8)
	c <- 1 + <-c
}
