package models

import (
	"github.com/opiumated/Conduit_2/config"
	"gopkg.in/mgo.v2/bson"
)

func BasicDataDb(udata User, config *config.Config) error {
	udata.Operation = registerOperation

	c := config.Database.C(userTableInfoString)
	err := c.Insert(udata)

	return err
}

func ProjectDataDb(uid string, projectData Project, config *config.Config) error {
	userCol := userProjectInfoString + uid
	c := config.Database.C(userCol)
	err := c.Insert(projectData)
	return err
}

func CoJoinProjectDb(jiskaJoinKrna string, coJoin CoJoinedProjectInfo, config *config.Config) error {
	jiskaJoinKrnaCollection := coJoinProjectInfoString + jiskaJoinKrna
	c := config.Database.C(jiskaJoinKrnaCollection)
	err := c.Insert(coJoin)
	return err
}

func JoinProDb(meraCol string, joinPro JoinedProjectInfo, config *config.Config) error {
	meraJoinedProject := joinedProjectInfoString + meraCol
	c := config.Database.C(meraJoinedProject)
	err := c.Insert(joinPro)
	return err
}

func LevDelCoJoinPro(juid, cuid, pid string, config *config.Config) error {
	coJoinCol := coJoinProjectInfoString + juid
	c := config.Database.C(coJoinCol)
	err := c.Remove(bson.M{checkPID: pid, checkCUID: cuid})
	return err
}

func LevDelJoinedProj(cuid, juid, pid string, config *config.Config) error {
	joinedPro := joinedProjectInfoString + cuid
	c := config.Database.C(joinedPro)
	err := c.Remove(bson.M{checkPID: pid, checkJUID: juid})
	return err
}

func DelFromMineProj(cuid, juid, pid string, config *config.Config) error {
	myOwnPro := userProjectInfoString + cuid
	c := config.Database.C(myOwnPro)
	err := c.Remove(bson.M{checkPID: pid})
	return err
}
