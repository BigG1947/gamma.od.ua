package main

import (
	"database/sql"
	"flag"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var router *mux.Router
var db *sql.DB

func main() {
	var runScripts bool
	flag.BoolVar(&runScripts, "init-db", false, "Run 'gamma.schema.sql' scripts for database")
	flag.Parse()

	var err error
	db, err = connectionToMysqlServer()
	if err != nil {
		log.Printf("ConnectionToMusqlServer: %s\n", err)
		return
	}

	if runScripts {
		err = createDbSchema()
		if err != nil {
			log.Printf("createDbSchema: %s\n", err)
			return
		}
	}

	router = routerInit()
	log.Fatal(http.ListenAndServe(":8080", router))
	return
}
