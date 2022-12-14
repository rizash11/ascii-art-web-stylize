package main

import (
	"html/template"
	"path/filepath"
)

type templateData struct {
	AsciiOutput string
}

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
