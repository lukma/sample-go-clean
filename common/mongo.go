package common

import (
	"os"

	mgo "gopkg.in/mgo.v2"
)

var session *mgo.Session

// GetConnection - Return database connection.
func GetConnection() (*mgo.Session, *mgo.Database) {
	if session == nil {
		s, err := mgo.Dial(os.Getenv("MONGO_HOST"))
		if err != nil {
			panic(err.Error())
		}
		session = s
	}

	s := session.Clone()
	return s, s.DB(os.Getenv("MONGO_DB"))
}
