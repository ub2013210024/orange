// Filename: cmd/api/routes.go
package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() *httprouter.Router {
	// Create a new router
	router := httprouter.New()
	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
	router.HandlerFunc(http.MethodPost, "/v1/schools", app.createSchoolHandler)
	router.HandlerFunc(http.MethodGet, "/v1/schools/:id", app.showSchoolHandler)
	// return the router
	return router

}
