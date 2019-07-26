package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api", getAPI)

	http.ListenAndServe(":80", router)
}

func getAPI(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello, world")
}
