package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func routerInit() *mux.Router {
	router := mux.NewRouter()

	// Main Routes
	router.HandleFunc("/", index)
	router.HandleFunc("/rss", rss)

	// News
	router.HandleFunc("/news", news)
	router.HandleFunc("/news/{id:[0-9]+}", singleNews)

	// Projects
	router.HandleFunc("/projects", projects)
	router.HandleFunc("/projects/{id:[0-9]+", singleProjects)

	// Video
	router.HandleFunc("/video", video)
	router.HandleFunc("/video/{id:[0-9]+", singleVideo)

	return router
}

func singleVideo(writer http.ResponseWriter, request *http.Request) {

}

func video(writer http.ResponseWriter, request *http.Request) {

}

func singleProjects(writer http.ResponseWriter, request *http.Request) {

}

func projects(writer http.ResponseWriter, request *http.Request) {

}

func singleNews(writer http.ResponseWriter, request *http.Request) {

}

func news(writer http.ResponseWriter, request *http.Request) {

}

func index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	_, err := fmt.Fprintf(w, "Index Page!\n")
	if err != nil {
		w.WriteHeader(500)
		log.Printf("%s\n", err)
		return
	}
	return
}

func rss(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	_, err := fmt.Fprintf(w, "About page!\n")
	if err != nil {
		w.WriteHeader(500)
		log.Printf("%s\n", err)
		return
	}
	return
}
