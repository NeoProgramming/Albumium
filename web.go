package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime/debug"
)

func InitWeb() {
	App.srv = &http.Server{
		Addr:     ":8081",
		ErrorLog: App.errorLog,
		Handler:  App.routes(),
	}
	fmt.Println("Web server initialized")
}

func HandleWeb() {
	err := App.srv.ListenAndServe()
	log.Fatal(err)
}

// The serverError helper writes an error message to the errorLog
// and then sends a 500 "Internal Server Error" response to the user.
func (app *Application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	//app.errorLog.Println(trace)
	app.errorLog.Output(2, trace)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

// The clientError helper sends a specific status code and corresponding description to user.
func (app *Application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

// It's just a convenience wrapper around clientError that sends a "404 Page Not Found" response to the user.
func (app *Application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}
