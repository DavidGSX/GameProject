package manager

import (
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var (
	session *mgo.Session
	dbName  = "test"
)

type KV struct {
	Key   string `bson:"_id"`
	Value string `bson:"v"`
}

func MongoDBInit() {
	if session == nil {
		var err error
		session, err = mgo.Dial("localhost")
		if err != nil {
			log.Panic(err)
		}
		session.SetMode(mgo.Monotonic, true)
		session.SetPoolLimit(30)
	}
}

func Session() *mgo.Session {
	if session == nil {
		MongoDBInit()
	}
	return session.Clone()
}

func mongoDBGetKV(t, k string) string {
	s := Session()
	defer s.Close()

	result := KV{}
	err := s.DB(dbName).C(t).Find(bson.M{"_id": k}).One(&result)
	if err != nil {
		return string("")
	}
	return result.Value
}

func mongoDBUpsertKV(t, k, v string) {
	s := Session()
	defer s.Close()

	s.DB(dbName).C(t).Upsert(bson.M{"_id": k}, KV{k, v})
}

func mongoDBDeleteKV(t, k string) {
	s := Session()
	defer s.Close()

	s.DB(dbName).C(t).Remove(bson.M{"_id": k})
}
