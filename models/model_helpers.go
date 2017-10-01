package models

import (
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func basicDataDb(udata User, session *mgo.Session) error {
	udata.Operation = registerOperation

	c := session.DB(databaseString).C(userTableInfoString)
	err := c.Insert(udata)

	return err
}

func projectDataDb(uid string, pdata ProjectInfo, session *mgo.Session) error {
	userCol := userProjectInfoString + uid
	c := session.DB(databaseString).C(userCol)
	err := c.Insert(pdata)

	return err
}

func coJoinProjectDb(jiskaJoinKrna string, coJoin CoJoinProjectInfo, session *mgo.Session) error {
	jiskaJoinKrnaCollection := coJoinProjectInfoString + jiskaJoinKrna
	c := session.DB(databaseString).C(jiskaJoinKrnaCollection)
	err := c.Insert(coJoin)

	return err
}

func joinProDb(meraCol string, joinPro JoinedProjectsInfo, session *mgo.Session) error {
	meraJoinedProject := joinedProjectInfoString + meraCol

	c := session.DB(databaseString).C(meraJoinedProject)
	err := c.Insert(joinPro)

	return err
}

func levDelCoJoinPro(juid, cuid, pid string, session *mgo.Session) error {
	coJoinCol := coJoinProjectInfoString + juid
	c := session.DB(databaseString).C(coJoinCol)

	err := c.Remove(bson.M{checkPID: pid, checkCUID: cuid})
	return err
}
func levDelJoinedProj(cuid, juid, pid string, session *mgo.Session) error {
	joinedPro := joinedProjectInfoString + cuid
	c := session.DB(databaseString).C(joinedPro)

	err := c.Remove(bson.M{checkPID: pid, checkJUID: juid})
	return err
}

func delFromMineProj(cuid, juid, pid string, session *mgo.Session) error {
	myOwnPro := userProjectInfoString + cuid
	c := session.DB(databaseString).C(myOwnPro)

	err := c.Remove(bson.M{checkPID: pid})
	return err
}

func checkIfRegistered(uid string, session *mgo.Session) error {
	result := User{}
	c := session.DB(databaseString).C(userTableInfoString)
	err := c.Find(bson.M{checkUID: uid}).One(&result)
	return err
}

func checkAndUpdate(udata User, session *mgo.Session) string {
	result := User{}

	c := session.DB(databaseString).C(userTableInfoString)
	err := c.Find(bson.M{checkUID: udata.UID, checkName: udata.Name,
		checkGender: udata.Gender, checkProfilePic: udata.ProfilePic}).One(&result)

	if err != nil {
		udata.Operation = updateProfileOperation

		colQuerier := bson.M{checkUID: udata.UID}
		err = c.Update(colQuerier, udata)

		return "Profile was updated"
	}

	return "No updates"
}
