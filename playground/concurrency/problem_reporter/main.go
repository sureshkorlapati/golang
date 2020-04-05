package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	restapi "github.com/sureshkorlapati/playground/concurrency/problem_reporter/rest"
)

func main() {

	port := 8080

	router := mux.NewRouter()

	router.Handle("/register", restapi.RestHandler{})

	err := http.ListenAndServe(":"+strconv.Itoa(port), router)
	if err != nil {
		log.Fatal("ListenAndServer error: ", err)
	}
}
