package handlers

import (
	"groupie-tracker/src/data"
	"html/template"
	"net/http"
)

type Error struct {
	Error   error
	Code    int
	Message string
}

var templates = template.Must(template.ParseGlob("src/templates/*.html"))

type Page struct {
	Title   string
	Error   string
	Artists []data.CombinedArtistData
}

func RenderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl, p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		ErrorHandler(w, r, http.StatusNotFound)
		return
	}
	p := &Page{Title: "Groupie Tracker"}
	RenderTemplate(w, "index.html", p)
}

func ArtistsHandler(w http.ResponseWriter, r *http.Request) {
	artists, err := data.FetchCombinedArtistData()
	if err != nil {
		http.Error(w, "Failed to fetch artist data", http.StatusInternalServerError)
		return
	}
	if r.URL.Path != "/artists" {
		ErrorHandler(w, r, http.StatusNotFound)
		return
	}
	p := &Page{Title: "Artists", Artists: artists}
	RenderTemplate(w, "artists.html", p)
}

func ErrorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	p := &Page{Title: "Error", Error: http.StatusText(status)}
	RenderTemplate(w, "error.html", p)
}

func NewError(err error, code int, msg string) *Error {
	return &Error{err, code, msg}
}

func ErrorPage(w http.ResponseWriter, r *http.Request) {
	p := &Page{Title: "Error", Error: "An unexpected error occurred."}
	RenderTemplate(w, "error.html", p)
}
