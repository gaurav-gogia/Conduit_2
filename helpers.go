package main

import (
	"log"

	mgo "gopkg.in/mgo.v2"
)

func logRegUpHelper(user userInformation) string {
	session, err := mgo.Dial(connectionString)
	if err != nil {
		log.Println(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	err = checkIfRegistered(user.UID, session)
	if err == nil {
		up := checkAndUpdate(user, session)
		log.Println(up)
		log.Println("Log In Success")

		return `{"response":"200"}`
	}

	err = basicDataDb(user, session)

	if err != nil {
		log.Println("Failed :/")
		return `{"response":"500"}`
	}

	log.Println("Sign Up Success")
	return `{"response":"200"}`
}

func addPojectHelper(uid string, pdata myProjectInfo) string {
	session, err := mgo.Dial(connectionString)
	if err != nil {
		log.Println(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	if err := projectDataDb(uid, pdata, session); err != nil {
		log.Println(err)
		return `{"response":"500"}`
	}

	log.Println("Project Added")
	return `{"response":"200"}`
}

func joinProjectHelper(joinPro joinedByMeProjectInfo, coJoin coJoinProjectInfo) string {
	session, err := mgo.Dial(connectionString)
	if err != nil {
		log.Println(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	err = coJoinProjectDb(joinPro.JUID, coJoin, session)
	err = joinProDb(coJoin.CUID, joinPro, session)

	if err != nil {
		log.Println(err)
		return `{"response":"500"}`
	}

	log.Println("Project Joined")
	return `{"response":"200"}`
}

func leaveProjectHelper(juid, cuid, pid string) string {
	session, err := mgo.Dial(connectionString)
	if err != nil {
		log.Println(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	err = levDelCoJoinPro(juid, cuid, pid, session)
	err = levDelJoinedProj(cuid, juid, pid, session)

	if err != nil {
		return `{"response":"500"}`
	}

	log.Println("Project left")
	return `{"response":"200"}`
}

func deleteProjectHelper(juid, cuid, pid string) string {
	session, err := mgo.Dial(connectionString)
	if err != nil {
		log.Println(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	err = levDelCoJoinPro(juid, cuid, pid, session)
	err = levDelJoinedProj(cuid, juid, pid, session)
	err = delFromMineProj(cuid, juid, pid, session)

	if err != nil {
		return `{"response":"500"}`
	}

	log.Println("Project left")
	return `{"response":"200"}`
}
