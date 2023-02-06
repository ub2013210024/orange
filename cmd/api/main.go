// Filename: cmd/api/main.go

package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// A Global variable to hold the applcation version number
const version = "1.0.0"

// Set up a struct ti hold server configuration
type config struct {
	port int
	env  string
}

// Set up dependency injection
type application struct {
	config config
	logger *log.Logger
}

// Set up main function
func main() {
	var cfg config
	// Get the arguments from the user or the server
	flag.IntVar(&cfg.port, "port", 4000, "API Server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment(development|staging|production server")
	flag.Parse()
	// Create a logger
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	// Create an object of type application
	app := &application{
		config: cfg,
		logger: logger,
	}
	// Create a route
	mux := http.NewServeMux()
	mux.HandleFunc("/v1/healthcheck", app.healthcheckHandler)

	// Create our server
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      mux,
		IdleTimeout:  time.Minute,      // inactive connection
		ReadTimeout:  10 * time.Second, // time to read request body or header
		WriteTimeout: 10 * time.Second,
	}

	logger.Printf("Starting %s on port %d", cfg.env, cfg.port)
	err := srv.ListenAndServe()
	log.Fatal(err)
}
