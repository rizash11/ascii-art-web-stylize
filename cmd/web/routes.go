package main

import "net/http"

// Keeps all of the handlers of the app.
func (app *application) routes() *http.ServeMux {
	router := http.NewServeMux()
	router.HandleFunc("/", app.home)
	router.HandleFunc("/ascii-art-web", app.asciiArtWeb)

	// Applies css to html templates.
	router.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./ui/static/"))))

	return router
}
