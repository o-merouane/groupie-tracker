package handlers

import "net/http"

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "src/templates/index.html")
}

// ErrorHandler serves the 404 page.
func ErrorHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "src/templates/404.html")
}

// CustomNotFoundHandler is a middleware that redirects to a 404 page if no content was found.
func CustomNotFoundHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rec := &statusRecorder{ResponseWriter: w, statusCode: http.StatusOK}
		next.ServeHTTP(rec, r)
		if rec.statusCode == http.StatusNotFound {
			http.Redirect(w, r, "/404", http.StatusSeeOther)
		}
	})
}

// statusRecorder is a wrapper around http.ResponseWriter that captures the status code.
type statusRecorder struct {
	http.ResponseWriter
	statusCode int
}

func (rec *statusRecorder) WriteHeader(statusCode int) {
	rec.statusCode = statusCode
	rec.ResponseWriter.WriteHeader(statusCode)
}
