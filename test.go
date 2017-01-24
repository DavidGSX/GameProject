package main

import (
	"fmt"
	"strconv"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Person struct {
	Name  string `bson:"_id"`
	Phone string `bson:"p"`
}

func main() {
	session, err := mgo.Dial("")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)
	session.SetPoolLimit(30)

	c := session.DB("test").C("people")
	for i := 1; i < 1e9; i++ {
		key := strconv.Itoa(i)
		begin := time.Now()
		c.Upsert(bson.M{"_id": key}, &Person{key, "111111"})

		result := Person{}
		err = c.Find(bson.M{"_id": key}).One(&result)
		if err != nil {
			panic(err)
		}
		if time.Since(begin) > 10*time.Millisecond {
			fmt.Println("Process", i, "	Use", time.Since(begin))
		}
	}
}
