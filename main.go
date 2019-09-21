package main

import (
	"database/sql"
	"flag"
	"github.com/gorilla/csrf"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"log"
	"net/http"
	"os"
	"time"
)

var router *mux.Router
var db *sql.DB
var store = sessions.NewCookieStore([]byte(os.Getenv("gamma_session_key")))

//var tokenV2 = "6Lf52rUUAAAAAC_l0PZJi6RiuJJwoKBXH4Clcfqi"
var tokenV3 = "6LfZbLEUAAAAAF5D1Xi2stbao2lPRYV-tItYDwGC"
var MAX_VALID_SCORE = 0.7
var location *time.Location

func main() {
	var runScripts bool
	flag.BoolVar(&runScripts, "init-db", false, "Run 'gamma.schema.sql' scripts for database")
	flag.Parse()

	var err error
	location, err = time.LoadLocation("Europe/Kiev")
	if err != nil {
		log.Printf("LoadLocation err: %s\n", err)
		return
	}

	db, err = connectionToMysqlServer()
	defer db.Close()
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

	CSRF := csrf.Protect(
		[]byte(os.Getenv("gamma_session_key")),
		csrf.FieldName("authenticity_token"),
		csrf.Secure(false),
		csrf.HttpOnly(false),
		csrf.Path("/"),
		csrf.MaxAge(0),
	)

	log.Fatal(http.ListenAndServe(":8080", CSRF(router)))
	return
}
