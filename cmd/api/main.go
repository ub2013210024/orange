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

// A global variable to hold the application
// version number
const version = "1.0.0"

// Setup a struct to hold server configuration
type config struct {
	port int
	env  string
}

// Setup dependency injection
type application struct {
	config config
	logger *log.Logger
}

// setup the main() function
func main() {
	var cfg config
	// Get the arguments for the user for
	// the server configuration
	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment(development|staging|production")
	flag.Parse()
	// Create a logger
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	// Create an object of type application
	app := &application{
		config: cfg,
		logger: logger,
	}
	// Create our server
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,      // inactive connections
		ReadTimeout:  10 * time.Second, // time to read request body or header
		WriteTimeout: 10 * time.Second,
	}
	// Start our server
	logger.Printf("starting %s server on port %d", cfg.env, cfg.port)
	err := srv.ListenAndServe()
	logger.Fatal(err)
}
