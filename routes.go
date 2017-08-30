package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	uuid "github.com/nu7hatch/gouuid"
)

func logRegUp(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {

		var user userInformation

		user.UID = r.FormValue("uid")
		user.Name = r.FormValue("uname")
		user.Gender = r.FormValue("usex")
		user.ProfilePic = r.FormValue("upic")
		user.TimeStamp = time.Now().Format(time.RFC850)

		response := logRegUpHelper(user)
		fmt.Fprintf(w, response)
	}

	log.Println(r.URL.Path)
}

func addProject(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var project myProjectInfo
		s, _ := uuid.NewV4()

		project.PID = s.String()
		project.Link = r.FormValue("link")
		project.ProjectName = r.FormValue("pname")
		project.ProjectDesc = r.FormValue("pdesc")
		project.TimeStamp = time.Now().Format(time.RFC850)
		uid := r.FormValue("uid")

		response := addPojectHelper(uid, project)

		fmt.Fprintf(w, response)
	}

	log.Println(r.URL.Path)
}

func joinProject(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var joinPro joinedByMeProjectInfo
		var coJoinPro coJoinProjectInfo

		joinPro.JUID = r.FormValue("juid")
		joinPro.PID = r.FormValue("pid")
		joinPro.ProjectName = r.FormValue("pname")

		coJoinPro.PID = joinPro.PID
		coJoinPro.CUID = r.FormValue("cuid")
		coJoinPro.CName = r.FormValue("cname")

		response := joinProjectHelper(joinPro, coJoinPro)

		fmt.Fprintf(w, response)
	}

	log.Println(r.URL.Path)
}

func leaveProject(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		pid := r.FormValue("pid")
		juid := r.FormValue("juid")
		cuid := r.FormValue("cuid")

		response := leaveProjectHelper(juid, cuid, pid)

		fmt.Fprintf(w, response)
	}

	log.Println(r.URL.Path)
}

func deleteProject(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		pid := r.FormValue("pid")
		juid := r.FormValue("juid")
		cuid := r.FormValue("cuid")

		response := leaveProjectHelper(juid, cuid, pid)

		fmt.Fprintf(w, response)
	}

	log.Println(r.URL.Path)
}
