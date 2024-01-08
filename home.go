package main

import (
	"html/template"
	"net/http"
)

type ViewHome struct {
	Title    string
	MainMenu int
}

func (app *Application) home(w http.ResponseWriter, r *http.Request) {
	data := ViewHome{}

	data.MainMenu = 0

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	// We initialize a slice containing paths to two files.
	files := []string{
		"./ui/pages/base.tmpl",
		"./ui/pages/home.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
		return
	}

	err = ts.Execute(w, data)
	if err != nil {
		app.serverError(w, err)
	}
}
