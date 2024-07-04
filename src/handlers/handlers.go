package handlers

import (
	"groupie-tracker/src/data"
	"html/template"
	"net/http"
)

type Page struct {
	Title   string
	Error   string
	Artists []data.CombinedArtistData
}

func RenderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	t, err := template.ParseFiles("src/templates/" + tmpl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := t.Execute(w, p); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	p := &Page{Title: "Groupie Tracker"}
	RenderTemplate(w, "index.html", p)
}

func ArtistsHandler(w http.ResponseWriter, r *http.Request) {
	artists, err := data.FetchCombinedArtistData()
	if err != nil {
		http.Error(w, "Failed to fetch artist data", http.StatusInternalServerError)
		return
	}
	p := &Page{Title: "Artists", Artists: artists}
	RenderTemplate(w, "artists.html", p)
}
