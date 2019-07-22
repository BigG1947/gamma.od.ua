package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func routerInit() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", index)

	return router
}

func index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	_, err := fmt.Fprintf(w, "Hello, world!")
	if err != nil {
		w.WriteHeader(500)
		log.Printf("%s\n", err)
		return
	}
	return
}
