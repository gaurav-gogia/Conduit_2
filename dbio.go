package main

import (
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func basicDataDb(udata userInformation, session *mgo.Session) error {
	udata.Operation = registerOperation

	c := session.DB(databaseString).C(userTableInfoString)
	err := c.Insert(udata)

	return err
}

func projectDataDb(uid string, pdata myProjectInfo, session *mgo.Session) error {
	userCol := userProjectInfoString + uid

	c := session.DB(databaseString).C(userCol)
	err := c.Insert(pdata)

	return err
}

func coJoinProjectDb(jiskaJoinKrna string, coJoin coJoinProjectInfo, session *mgo.Session) error {
	jiskaJoinKrnaCollection := coJoinProjectInfoString + jiskaJoinKrna

	c := session.DB(databaseString).C(jiskaJoinKrnaCollection)
	err := c.Insert(coJoin)

	return err
}

func joinProDb(meraCol string, joinPro joinedByMeProjectInfo, session *mgo.Session) error {
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
