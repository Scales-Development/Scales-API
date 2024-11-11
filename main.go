package main

import (
	"fmt"
	"net/http"
	"scares/scales-api/login"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/login", login.LoginHandler).Methods("POST")
	router.HandleFunc("/protected", login.ProtectedHandler).Methods("GET")

	fmt.Println("Starting up the Scales API....")
	err := http.ListenAndServe("localhost:4000", router)
	if err != nil {
		fmt.Println("An Unexpected error has occurred while starting up the API.", err)
	}
	fmt.Println("Server API Started. Listening on port 4000")
}
