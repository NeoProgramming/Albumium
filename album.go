package main

import (
	"html/template"
	"net/http"
)

type Pagination struct {
	Count      int
	Page       int
	TotalPages int
	PrevPage   int
	NextPage   int
}

type ViewAlbum struct {
	Pagination
	Title string
	MainMenu int
}

func (app *Application) album(w http.ResponseWriter, r *http.Request) {
	data := ViewAlbum{
	}

	// We initialize a slice containing paths to two files.
	files := []string{
		"./ui/pages/base.tmpl",
		"./ui/pages/album.tmpl",
		"./ui/fragments/pagination.tmpl",
		"./ui/fragments/media.tmpl",
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
