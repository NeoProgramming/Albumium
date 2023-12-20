package main

import (
	"fmt"
	"net/http"
)

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./ui/static/images/favicon.ico")
}

func (app *Application) routes() *http.ServeMux {

	fmt.Println("routes init")

	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static"))
	mediaServer := http.FileServer(http.Dir(App.config.BasePath))
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))
	mux.Handle("/media/", http.StripPrefix("/media/", mediaServer))
	mux.HandleFunc("/favicon.ico", faviconHandler)

	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/album", app.album)
	//	mux.HandleFunc("/groups", app.groups)

	return mux
}
