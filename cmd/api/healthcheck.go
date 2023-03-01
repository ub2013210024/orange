package main

import (
	"fmt"
	"net/http"
)

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "status: available\n")
	// fmt.Fprintf(w, "env: %s\n", app.config.env)
	// fmt.Fprintf(w, "version: %s\n", version)
	js := `{"status":"available", "environment":%q, "version": %q}`
	js = fmt.Sprintf(js, app.config.env, version)
	// add header information
	w.Header().Set("Content-Type", "application/json")
}
