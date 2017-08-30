package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/logRegUp", logRegUp)
	http.HandleFunc("/addProject", addProject)
	http.HandleFunc("/joinProject", joinProject)
	http.HandleFunc("/leaveProject", leaveProject)
	http.HandleFunc("/deleteProject", deleteProject)

	fmt.Println("Server listening at ", port)
	http.ListenAndServe(port, nil)
}
