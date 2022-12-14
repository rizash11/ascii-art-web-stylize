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
	address := flag.String("address", "4000", "web HTTP address")
	flag.Parse()

	errorLog := log.New(os.Stderr, "ERROR: \t", log.Ldate|log.Ltime|log.Lshortfile)
	infoLog := log.New(os.Stdout, "INFO \t", log.Ldate|log.Ltime)

	templateCache, err := newTemplateCache("ui/html")
	if err != nil {
		errorLog.Fatalln(err)
	}

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
