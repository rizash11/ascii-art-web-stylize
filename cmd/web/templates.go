package main

import (
	"html/template"
	"path/filepath"
)

// A single structure for the data that is to be sent to html templates.
// Here only one variable of data is used, but more can be added if necessary.
type templateData struct {
	AsciiOutput string
}

// Takes html files in a given directory, parses them, and stores in a map
func newTemplateCache(dir string) (map[string]*template.Template, error) {
	templateCache := map[string]*template.Template{}
	pages, err := filepath.Glob(filepath.Join(dir, "*.page.html"))
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		ts, err := template.ParseFiles(page)
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseGlob(filepath.Join(dir, "*layout.html"))
		if err != nil {
			return nil, err
		}

		name := filepath.Base(page)
		templateCache[name] = ts
	}

	return templateCache, nil
}
