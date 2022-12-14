package main

import (
	"fmt"
	"net/http"
	"runtime/debug"
)

func (app *application) notFound(w http.ResponseWriter) {
	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
}

func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errorLog.Output(2, trace)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *application) render(w http.ResponseWriter, r *http.Request, name string, td *templateData) {
	tmpl, ok := app.templateCache[name]
	if !ok {
		app.serverError(w, fmt.Errorf("the %s page doesn't exist", name))
		return
	}

	err := tmpl.Execute(w, td)
	if err != nil {
		app.serverError(w, err)
		return
	}
}
