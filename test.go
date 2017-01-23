package main

import (
	"fmt"
	"strconv"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Person struct {
	Phone string
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
		err = c.Insert(bson.M{"_id": key, "Phone": "111111"}) //(&Person{"Ale", "111111"}, &Person{"Cla", "222222222"})
		if err != nil {
			panic(err)
		}
		result := bson.M{}
		err = c.Find(bson.M{"_id": key}).One(result)
		if err != nil {
			panic(err)
		}
		if time.Since(begin) > 50*time.Millisecond {
			fmt.Println("Process --->>", i, time.Since(begin))
		}
	}
}
