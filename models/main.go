package models

import (
	"log"

	"github.com/DesmondANIMUS/Conduit_2/config"

	mgo "gopkg.in/mgo.v2"
)

//LogRegUpHelper Authenticates a user
func LogRegUpHelper(cfg *config.Config, user User) string {
	session := cfg.Session.Copy()
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	err := checkIfRegistered(user.UID, session)
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

func AddPojectHelper(cfg *config.Config, uid string, pdata ProjectInfo) string {
	session := cfg.Session.Copy()
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	if err := projectDataDb(uid, pdata, session); err != nil {
		log.Println(err)
		return `{"response":"500"}`
	}

	log.Println("Project Added")
	return `{"response":"200"}`
}

func JoinProjectHelper(cfg *config.Config, joinPro JoinedProjectsInfo, coJoin CoJoinProjectInfo) string {
	session := cfg.Session.Copy()
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	err := coJoinProjectDb(joinPro.JUID, coJoin, session)
	err = joinProDb(coJoin.CUID, joinPro, session)

	if err != nil {
		log.Println(err)
		return `{"response":"500"}`
	}

	log.Println("Project Joined")
	return `{"response":"200"}`
}

func LeaveProjectHelper(cfg *config.Config, juid, cuid, pid string) string {
	session := cfg.Session.Copy()
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	err := levDelCoJoinPro(juid, cuid, pid, session)
	err = levDelJoinedProj(cuid, juid, pid, session)

	if err != nil {
		return `{"response":"500"}`
	}

	log.Println("Project left")
	return `{"response":"200"}`
}

func DeleteProjectHelper(cfg *config.Config, juid, cuid, pid string) string {
	session := cfg.Session.Copy()
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	err := levDelCoJoinPro(juid, cuid, pid, session)
	err = levDelJoinedProj(cuid, juid, pid, session)
	err = delFromMineProj(cuid, juid, pid, session)

	if err != nil {
		return `{"response":"500"}`
	}

	log.Println("Project left")
	return `{"response":"200"}`
}
