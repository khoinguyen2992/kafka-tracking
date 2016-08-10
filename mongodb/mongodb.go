package mongodb

import (
	"os"
	"strings"
	"time"

	mgo "gopkg.in/mgo.v2"
)

var Session *mgo.Session

func init() {
	if Session == nil {
		Session = getConnection()
	}
}

func getConnection() *mgo.Session {
	s, err := mgo.DialWithInfo(&mgo.DialInfo{
		Addrs:    strings.Split(os.Getenv("MONGODB_HOST"), ","),
		Timeout:  5 * time.Second,
		Database: os.Getenv("MONGODB_DATABASE"),
		Username: os.Getenv("MONGODB_USERNAME"),
		Password: os.Getenv("MONGODB_PASSWORD"),
	})

	if err != nil {
		panic(err)
	}

	return s
}
