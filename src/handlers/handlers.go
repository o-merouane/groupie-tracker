package handlers

import (
	"encoding/json"
	"groupie-tracker/src/data"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "src/templates/index.html")
}

func ArtistsHandler(w http.ResponseWriter, r *http.Request) {
	artists, err := data.FetchArtistData()
	if err != nil {
		http.Error(w, "Failed to fetch artist data", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(artists)
}
