package handlers

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/DesmondANIMUS/Conduit_2/config"
	"github.com/DesmondANIMUS/Conduit_2/models"
	uuid "github.com/nu7hatch/gouuid"
)

func LogRegUp(cfg *config.Config) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		var user models.User
		user.UID = r.FormValue("uid")
		user.Name = r.FormValue("uname")
		user.Gender = r.FormValue("usex")
		user.ProfilePic = r.FormValue("upic")
		user.TimeStamp = time.Now().Format(time.RFC850)

		response := models.LogRegUpHelper(cfg, user)
		fmt.Fprintf(w, response)

		log.Println(r.URL.Path)
	}
	return http.HandlerFunc(fn)
}

func AddProject(cfg *config.Config) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		var project models.ProjectInfo
		s, _ := uuid.NewV4()
		project.PID = s.String()
		project.Link = r.FormValue("link")
		project.ProjectName = r.FormValue("pname")
		project.ProjectDesc = r.FormValue("pdesc")
		project.TimeStamp = time.Now().Format(time.RFC850)
		uid := r.FormValue("uid")

		response := models.AddPojectHelper(cfg, uid, project)

		fmt.Fprintf(w, response)
		log.Println(r.URL.Path)
	}
	return http.HandlerFunc(fn)
}

func JoinProject(cfg *config.Config) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		var joinPro models.JoinedProjectsInfo
		var coJoinPro models.CoJoinProjectInfo

		joinPro.JUID = r.FormValue("juid")
		joinPro.PID = r.FormValue("pid")
		joinPro.ProjectName = r.FormValue("pname")

		coJoinPro.PID = joinPro.PID
		coJoinPro.CUID = r.FormValue("cuid")
		coJoinPro.CName = r.FormValue("cname")

		response := models.JoinProjectHelper(cfg, joinPro, coJoinPro)

		fmt.Fprintf(w, response)
		log.Println(r.URL.Path)
	}

	return http.HandlerFunc(fn)
}

func LeaveProject(cfg *config.Config) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		pid := r.FormValue("pid")
		juid := r.FormValue("juid")
		cuid := r.FormValue("cuid")

		response := models.LeaveProjectHelper(cfg, juid, cuid, pid)

		fmt.Fprintf(w, response)
		log.Println(r.URL.Path)
	}
	return http.HandlerFunc(fn)
}

func DeleteProject(cfg *config.Config) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		pid := r.FormValue("pid")
		juid := r.FormValue("juid")
		cuid := r.FormValue("cuid")
		response := models.LeaveProjectHelper(cfg, juid, cuid, pid)
		fmt.Fprintf(w, response)
		log.Println(r.URL.Path)
	}
	return http.HandlerFunc(fn)
}
