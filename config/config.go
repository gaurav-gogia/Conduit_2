package config

import (
	"log"
	"os"

	mgo "gopkg.in/mgo.v2"
)

//Config Represents the configuration of our app
//Add more here
type Config struct {
	MongoServer string
	MongoDB     string
	Port        string
	Database    *mgo.Database
}

//Init Initializes tbe Config
func Init() {
	MONGOSERVER := os.Getenv("MONGO_URL")
	MONGODB := os.Getenv("MONGO_DB")
	if MONGOSERVER == "" {
		log.Println("No specified Mongo Address, using default")
		MONGOSERVER = "mongodb://localhost/"
		MONGODB = "ConduitDB"
	}

	session, err := mgo.Dial(MONGOSERVER)
	if err != nil {
		log.Println(err)
	}
	session.SetMode(mgo.Monotonic, true)
	defer session.Close()
	config := Config{
		MongoServer: MONGOSERVER,
		MongoDB:     MONGODB,
		Database:    session.DB(MONGODB),
		Port:        "8080",
	}
	log.Println("Setup Configuration: ", config)
}

func Get() *Config {
	return &Config{}
}
