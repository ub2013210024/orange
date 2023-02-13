// Filename: cmd/api/schools.go

package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (app *application) createSchoolHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Created a new school")
}

func (app *application) showSchoolHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "School displayed")
	params := httprouter.ParamsFromContext(r.Context())
	id, err := strconv.ParseInt(params.ByName("id"), 10, 10)
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
}
