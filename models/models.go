package models

const (
	port                    = ":8888"
	databaseString          = "ConduitDB"
	connectionString        = "mongodb://localhost/"
	userTableInfoString     = "userInfo"
	userProjectInfoString   = "userProj_"
	joinedProjectInfoString = "joinedProj_"
	coJoinProjectInfoString = "coJoinProj_"

	checkUID        = "uid"
	checkGender     = "gender"
	checkPID        = "pid"
	checkJUID       = "juid"
	checkCUID       = "cuid"
	checkProfilePic = "profile_pic"
	checkName       = "name"

	registerOperation      = "REGISTER"
	updateProfileOperation = "UPDATE_PROFILE"
)

//User Represents a user in our collection
type User struct {
	UID        string `bson:"uid"`
	Name       string `bson:"name"`
	Gender     string `bson:"gender"` // 0 = M, 1 = F, 2 = Other
	ProfilePic string `bson:"profile_pic"`
	TimeStamp  string `bson:"time"`
	Operation  string `bson:"operation"`
}

type ProjectInfo struct {
	PID         string `bson:"pid"`
	ProjectName string `bson:"project_name"`
	ProjectDesc string `bson:"project_desc"`
	Link        string `bson:"link"`
	TimeStamp   string `bson:"time"`
}

type JoinedProjectsInfo struct {
	JUID        string `bson:"juid"`         // jiska project uski uid to find their collection
	PID         string `bson:"pid"`          // jiska project join kra uski pid
	ProjectName string `bson:"project_name"` // jiska project join kra uske project ka name
}

type CoJoinProjectInfo struct {
	PID   string `bson:"pid"`   // jis project pe vo kaam kr rha uski pid
	CUID  string `bson:"cuid"`  // jisne mera project join kra uski uid
	CName string `bson:"cname"` // jisne mera project join kra uska name
}
