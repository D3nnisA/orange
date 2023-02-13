// Filename: cmd/api/schools.go

package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (app *application) createSchoolHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "created a new school...")
}

func (app *application) showSchoolHandler(w http.ResponseWriter, r *http.Request) {

	params := httprouter.ParamsFromContext(r.Context())
	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)

	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "show details of school %d\n", id)

}
