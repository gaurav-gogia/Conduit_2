package helper

import (
	"log"

	"github.com/opiumated/Conduit_2/config"
	mgo "gopkg.in/mgo.v2"
)

func LogRegUpHelper(user User, config *config.Config) string {
	session := config.Database.Session.Copy()
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	err := models.CheckIfRegistered(user.UID, config)
	if err == nil {
		up := checkAndUpdate(user, config)
		log.Println(up)
		log.Println("Log In Success")

		return `{"response":"200"}`
	}

	err = models.BasicDataDb(user, config)

	if err != nil {
		log.Println("Failed :/")
		return `{"response":"500"}`
	}

	log.Println("Sign Up Success")
	return `{"response":"200"}`
}

func AddPojectHelper(uid string, pdata Project, config *config.Config) string {
	session := config.Database.Session.Copy()
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	if err := models.ProjectDataDb(uid, pdata, config); err != nil {
		log.Println(err)
		return `{"response":"500"}`
	}

	log.Println("Project Added")
	return `{"response":"200"}`
}

func JoinProjectHelper(joinPro JoinedProjectInfo, coJoin CoJoinedProjectInfo, config *config.Config) string {
	session := config.Database.Session.Copy()
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	err := models.CoJoinProjectDb(joinPro.JUID, coJoin, config)
	err = JoinProDb(coJoin.CUID, joinPro, config)

	if err != nil {
		log.Println(err)
		return `{"response":"500"}`
	}
	log.Println("Project Joined")
	return `{"response":"200"}`
}

func LeaveProjectHelper(juid, cuid, pid string, config *config.Config) string {
	session := config.Database.Session.Copy()
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	err := models.LevDelCoJoinPro(juid, cuid, pid, config)
	err = models.LevDelJoinedProj(cuid, juid, pid, config)

	if err != nil {
		return `{"response":"500"}`
	}

	log.Println("Project left")
	return `{"response":"200"}`
}

func DeleteProjectHelper(juid, cuid, pid string, config *config.Config) string {
	session := config.Database.Session.Copy()
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	err := models.LevDelCoJoinPro(juid, cuid, pid, config)
	err = models.LevDelJoinedProj(cuid, juid, pid, config)
	err = models.DelFromMineProj(cuid, juid, pid, config)

	if err != nil {
		return `{"response":"500"}`
	}

	log.Println("Project left")
	return `{"response":"200"}`
}
