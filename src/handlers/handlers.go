package handlers

import (
	"encoding/json"
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

func LocationsHandler(w http.ResponseWriter, r *http.Request) {
	url := "https://groupietrackers.herokuapp.com/api/locations"

	locations, err := data.FetchLocationData(url)
	if err != nil {
		http.Error(w, "Failed to fetch locations data", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(locations); err != nil {
		http.Error(w, "Failed to encode locations data", http.StatusInternalServerError)
		return
	}
}

func DatesHandler(w http.ResponseWriter, r *http.Request) {
	url := "https://groupietrackers.herokuapp.com/api/dates"
	dates, err := data.FetchDatesData(url)
	if err != nil {
		http.Error(w, "Failed to fetch dates data", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(dates); err != nil {
		http.Error(w, "Failed to encode dates data", http.StatusInternalServerError)
		return
	}
}

func RelationsHandler(w http.ResponseWriter, r *http.Request) {
	url := "https://groupietrackers.herokuapp.com/api/relation"

	relations, err := data.FetchRelationsData(url)
	if err != nil {
		http.Error(w, "Failed to fetch relations data", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(relations); err != nil {
		http.Error(w, "Failed to encode relations data", http.StatusInternalServerError)
		return
	}
}
