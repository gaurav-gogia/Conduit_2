package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	config2 "github.com/DesmondANIMUS/Conduit_2/config"
	"github.com/DesmondANIMUS/Conduit_2/handlers"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	config := config2.InitConfig()

	router := mux.NewRouter()

	router.Handle("/logRegUp", handlers.LogRegUp(config)).Methods(http.MethodPost)
	router.Handle("/addProject", handlers.AddProject(config))
	router.Handle("/joinProject", handlers.JoinProject(config))
	router.Handle("/leaveProject", handlers.LeaveProject(config))
	router.Handle("/deleteProject", handlers.DeleteProject(config))

	server := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:3000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(server.ListenAndServe())
}
