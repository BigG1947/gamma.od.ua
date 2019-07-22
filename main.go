package main

import (
	"database/sql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var router *mux.Router
var db *sql.DB

func main() {

	router = routerInit()
	log.Fatal(http.ListenAndServe(":8080", router))
	return
}
