package cacheMgr

import (
	"log"
	"strings"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type KV struct {
	Key   string `bson:"_id"`
	Value string `bson:"v"`
}

var (
	session *mgo.Session
	dbName  = "test"
)

func mongoDBInit() {
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
		mongoDBInit()
	}
	return session.Clone()
}

func getTK(key string) (t, k string) {
	tk := strings.Split(key, "_")
	if len(tk) < 2 {
		log.Panic("key invalid ", key)
	}
	t = tk[0]
	k = tk[1]
	for i := 2; i < len(tk); i++ {
		k = k + "_" + tk[i]
	}
	return t, k
}

func mongoDBGetKV(key string) string {
	s := Session()
	defer s.Close()

	t, k := getTK(key)
	result := KV{}
	err := s.DB(dbName).C(t).Find(bson.M{"_id": k}).One(&result)
	if err != nil {
		return string("")
	}
	return result.Value
}

func mongoDBUpsertKV(key, v string) {
	s := Session()
	defer s.Close()

	t, k := getTK(key)
	s.DB(dbName).C(t).Upsert(bson.M{"_id": k}, KV{k, v})
}

func mongoDBDeleteKV(key string) {
	s := Session()
	defer s.Close()

	t, k := getTK(key)
	s.DB(dbName).C(t).Remove(bson.M{"_id": k})
}
