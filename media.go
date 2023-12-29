package main

import (
	"fmt"
	"os/exec"
	"net/http"
	"strconv"
	"strings"
)

func (app *Application) openMedia(w http.ResponseWriter, r *http.Request) {
	id := Atoi(r.URL.Path[len("/open-media/"):])
	fmt.Println("media id = ", id)
	// select from db
	m := getMediaById(App.db, id)
	if m == nil {
		fmt.Println("openMedia m=nil")
		return
	}
	
	// open in external program
	ft := FileType(m.Path)
	ap := FileApp(ft)
	cmd := exec.Command(ap, m.Path)
	
	err := cmd.Run()
	if err != nil {
		fmt.Println("Exec.Run error:", err)
	}
}

func extractCheckboxes(w http.ResponseWriter, r *http.Request) []int {
	// Parse the form data
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return make([]int, 0)
	}

	// Get the values of the checkboxes
	checkboxValues := r.FormValue("checkbox")
	fmt.Println(checkboxValues)

	// Split the comma-separated values into a slice
	values := strings.Split(checkboxValues, ",")

	ints := make([]int, 0, len(values))
	for _, token := range values {
		if i, err := strconv.Atoi(token); err == nil {
			ints = append(ints, i)
		}
	}
	return ints
}

func (app *Application) handleMedia(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handleMedia")
	ids := extractCheckboxes(w, r)
	for _, id := range ids {
		fmt.Println("check ", id)
	}
}


