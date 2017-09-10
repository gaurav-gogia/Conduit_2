package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	uuid "github.com/nu7hatch/gouuid"
	"github.com/opiumated/Conduit_2/models"
)

func LogRegUpHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {

		var user models.User

		user.UID = r.FormValue("uid")
		user.Name = r.FormValue("uname")
		user.Gender = r.FormValue("usex")
		user.ProfilePic = r.FormValue("upic")
		user.TimeStamp = time.Now().Format(time.RFC850)
		response := models.LogRegUpHelper(user)
		fmt.Fprintf(w, response)
	}

	log.Println(r.URL.Path)
}

func AddProjectHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var project myProjectInfo
		s, _ := uuid.NewV4()

		project.PID = s.String()
		project.Link = r.FormValue("link")
		project.ProjectName = r.FormValue("pname")
		project.ProjectDesc = r.FormValue("pdesc")
		project.TimeStamp = time.Now().Format(time.RFC850)
		uid := r.FormValue("uid")

		response := models.AddPojectHelper(uid, project)

		fmt.Fprintf(w, response)
	}

	log.Println(r.URL.Path)
}

func JoinProjectHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var joinPro joinedByMeProjectInfo
		var coJoinPro coJoinProjectInfo

		joinPro.JUID = r.FormValue("juid")
		joinPro.PID = r.FormValue("pid")
		joinPro.ProjectName = r.FormValue("pname")

		coJoinPro.PID = joinPro.PID
		coJoinPro.CUID = r.FormValue("cuid")
		coJoinPro.CName = r.FormValue("cname")

		response := models.JoinProjectHelper(joinPro, coJoinPro)

		fmt.Fprintf(w, response)
	}

	log.Println(r.URL.Path)
}

func LeaveProjectHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		pid := r.FormValue("pid")
		juid := r.FormValue("juid")
		cuid := r.FormValue("cuid")

		response := models.LeaveProjectHelper(juid, cuid, pid)

		fmt.Fprintf(w, response)
	}

	log.Println(r.URL.Path)
}

func DeleteProjectHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		pid := r.FormValue("pid")
		juid := r.FormValue("juid")
		cuid := r.FormValue("cuid")

		response := models.LeaveProjectHelper(juid, cuid, pid)

		fmt.Fprintf(w, response)
	}

	log.Println(r.URL.Path)
}
