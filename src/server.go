package src

import (
	"groupie-tracker/src/handlers"
	"log"
	"net/http"
)

func Server() {
	fs := http.FileServer(http.Dir("src/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", handlers.ArtistsHandler)
	http.HandleFunc("/error", handlers.ErrorPage)

	log.Println("Server started on: http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
