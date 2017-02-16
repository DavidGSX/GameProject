package main

import (
	"fmt"
	//"math/rand"
	"runtime"
	"strconv"
	"time"

	"gopkg.in/mgo.v2"
	//"gopkg.in/mgo.v2/bson"
)

type Person struct {
	Name  string `bson:"_id"`
	Phone string `bson:"value"`
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	for begin := 0; begin < 2e9; begin += 1e8 {
		go upsert(begin)
	}
	<-time.After(1e9 * time.Second)
}

func upsert(begin int) {
	session, err := mgo.Dial("10.137.17.59")
	if err != nil {
		fmt.Println(err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)
	session.SetPoolLimit(5)

	var i int
	c := session.DB("test").C("GoodPeople")
	for i = begin; i < begin+1e8; i++ {
		key := strconv.Itoa(i) //rand.Intn(510000) + begin)

		value := key
		for len(value) < 1000 {
			value += "1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890"
		}
		c.Insert(Person{key, value})

		//c.Remove(bson.M{"_id": key})
		/*
			result := Person{}
			err := c.Find(bson.M{"_id": key}).One(&result)
			if err != nil {
				//fmt.Println(err)
			}
		*/
		if (i-begin)%10000 == 0 {
			fmt.Println(begin, "------Process------", (i - begin))
		}
	}
}
