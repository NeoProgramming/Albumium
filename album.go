package main

import (
	"math"
	"net/http"
	"strconv"
	"text/template"
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
	PageExtraArg string // to remove?
	Title        string
	MainMenu     int
	Files        []Media
}

func Atodi(s string, d int) int {
	if i, err := strconv.Atoi(s); err == nil {
		return i
	}
	return d
}

func (app *Application) album(w http.ResponseWriter, r *http.Request) {
	data := ViewAlbum{}

	// extract args
	data.Count = getMediaCount(App.db)
	data.Page = Atodi(r.URL.Query().Get("page"), 1)
	data.NextPage = data.Page + 1
	data.PrevPage = data.Page - 1
	data.TotalPages = int(math.Ceil(float64(data.Count) / float64(25)))

	// select files
	data.Files = getMedia(App.db, data.Page, 25, "", "", "", false)

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
