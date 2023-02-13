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

	fmt.Fprintln(w, "School display...")
}
