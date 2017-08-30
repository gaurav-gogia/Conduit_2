package main

import (
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func checkIfRegistered(uid string, session *mgo.Session) error {
	result := userInformation{}

	c := session.DB(databaseString).C(userTableInfoString)
	err := c.Find(bson.M{checkUID: uid}).One(&result)

	return err
}
func checkAndUpdate(udata userInformation, session *mgo.Session) string {
	result := userInformation{}

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
