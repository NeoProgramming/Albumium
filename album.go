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
	PageExtraArg string
	Title        string
	MainMenu     int
	Files        []Media
	Filters      string
	Search       string
	FType        int
}

func (app *Application) album(w http.ResponseWriter, r *http.Request) {
	data := ViewAlbum{}

	// extract args
	data.MainMenu = 1
	data.Search = r.URL.Query().Get("search")
	data.FType = Atodi(r.URL.Query().Get("ftype"), 0)
	data.Filters = r.URL.Query().Get("filters")

	data.Count = getMediaCount(App.db, data.Search, data.FType, data.Filters)
	data.Page = Atodi(r.URL.Query().Get("page"), 1)
	data.NextPage = data.Page + 1
	data.PrevPage = data.Page - 1
	data.TotalPages = int(math.Ceil(float64(data.Count) / float64(25)))

	if data.Search != "" {
		data.PageExtraArg += "&search=" + data.Search
	}
	if data.FType != 0 {
		data.PageExtraArg += "&ftype=" + strconv.Itoa(data.FType)
	}

	// select files
	data.Files = getMedia(App.db, data.Page, 25, data.Search, data.FType, "")

	// We initialize a slice containing paths to two files.
	files := []string{
		"./ui/pages/base.tmpl",
		"./ui/pages/album.tmpl",
		"./ui/fragments/pagination.tmpl",
		"./ui/fragments/media.tmpl",
		"./ui/fragments/search.tmpl",
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
