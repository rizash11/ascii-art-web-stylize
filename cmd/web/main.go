package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"
)

type application struct {
	errorLog      *log.Logger
	infoLog       *log.Logger
	templateCache map[string]*template.Template
}

func main() {
	// This flag lets a user choose a port for the app.
	address := flag.String("address", "4000", "web HTTP address")
	flag.Parse()

	// These two loggers will document occuring errors and other information,
	// and they will all be available within one structure that is going to be created for the app.
	errorLog := log.New(os.Stderr, "ERROR: \t", log.Ldate|log.Ltime|log.Lshortfile)
	infoLog := log.New(os.Stdout, "INFO \t", log.Ldate|log.Ltime)

	// This variable keeps all of html templates parsed and ready to execute,
	// so that the templates are not parsed repeatedly each time they are used.
	templateCache, err := newTemplateCache("ui/html")
	if err != nil {
		errorLog.Fatalln(err)
	}

	// Creates a single variable that keeps all the logs and templates of the app
	asciiArtWeb := application{
		errorLog:      errorLog,
		infoLog:       infoLog,
		templateCache: templateCache,
	}

	srv := &http.Server{
		Addr:    "localhost:" + *address,
		Handler: asciiArtWeb.routes(),
	}

	infoLog.Println("Starting a server at http://localhost:" + *address + "/")
	errorLog.Fatalln(srv.ListenAndServe())
}
