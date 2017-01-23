package cacheMgr

import (
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

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

func mongoDBGetKV(t, k string) string {
	s := Session()
	defer s.Close()

	result := bson.M{}
	err := s.DB(dbName).C(t).Find(bson.M{"_id": k}).One(&result)
	if err != nil {
		return string("")
	}
	v, ok := result["v"].(string)
	if ok {
		return v
	} else {
		log.Panic("found but nil", k, result)
		return string("")
	}
}

func mongoDBInsertKV(t, k, v string) {
	s := Session()
	defer s.Close()

	s.DB(dbName).C(t).Insert(bson.M{"_id": k, "v": v})
}

func mongoDBUpdateKV(t, k, v string) {
	s := Session()
	defer s.Close()

	s.DB(dbName).C(t).Update(bson.M{"_id": k}, bson.M{"_id": k, "v": v})
}

func mongoDBDeleteKV(t, k string) {
	s := Session()
	defer s.Close()

	s.DB(dbName).C(t).Remove(bson.M{"_id": k})
}
