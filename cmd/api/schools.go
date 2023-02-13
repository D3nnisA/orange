// Filename: cmd/api/schools.go

package main

import (
	"fmt"
	"net/http"
)

func (app *application) createSchoolHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "created a new school...")
}

func (app *application) showSchoolHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParams(r)

	if err != nil {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "show details of school %d\n", id)

}
