package models

import (
	"github.com/opiumated/Conduit_2/config"
	"gopkg.in/mgo.v2/bson"
)

func checkIfRegistered(uid string, config *config.Config) error {
	result := User{}
	c := config.Database.C(userTableInfoString)
	err := c.Find(bson.M{checkUID: uid}).One(&result)
	return err
}
func checkAndUpdate(udata User, config *config.Config) string {
	result := User{}
	c := config.Database.C(userTableInfoString)
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
