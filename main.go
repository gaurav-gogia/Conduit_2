package main

import (
	"fmt"
	"net/http"

	"github.com/opiumated/Conduit_2/config"
)

const (
	PORT = "8080"
)

func main() {
	config.Init()

	http.HandleFunc("/logRegUp", logRegUp)
	http.HandleFunc("/addProject", addProject)
	http.HandleFunc("/joinProject", joinProject)
	http.HandleFunc("/leaveProject", leaveProject)
	http.HandleFunc("/deleteProject", deleteProject)

	fmt.Println("Server listening at ", PORT)
	http.ListenAndServe(PORT, nil)
}
