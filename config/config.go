package config

import (
	"gopkg.in/mgo.v2"
	"log"
	"os"
)

type Config struct {
	MongoServer string
	MongoDB string
	Session *mgo.Session
	Database *mgo.Database
}


func InitConfig() *Config{
	MONGOSERVER := os.Getenv("MONGO_SERVER")
	MONGODB := os.Getenv("MONGO_DB")
	if MONGOSERVER == "" {
		MONGOSERVER = "localhost:27017"
		MONGODB = "conduit2"
	}
	session, err := mgo.Dial(MONGOSERVER)
	if err != nil {
		log.Fatal(err)
	}

	config := &Config{
		MongoServer: MONGOSERVER,
		MongoDB: MONGODB,
		Session: session,
		Database: session.DB(MONGODB),
	}

	return config
}

func getConfig() *Config {
	return &Config{}
}