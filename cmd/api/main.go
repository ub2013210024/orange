// Filename: cmd/api/main.go

package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

// A Global variable to hold the applcation version number
const version = "1.0.0"

// Set up a struct ti hold server configuration
type config struct {
	port int
	env  string
	db   struct {
		dsn string
	}
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
	flag.StringVar(&cfg.db.dsn, "dsn", os.Getenv("LEMON_DB_DSN"), "PostgresSQL DSN")
	flag.Parse()
	// Create a logger
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	// Setup the database connection pool
	db, err := openDB(cfg)
	if err != nil {
		logger.Fatal(err)
	}
	// Close the database connection pool
	defer db.Close()

	// Create an object of type application
	app := &application{
		config: cfg,
		logger: logger,
	}
	// Create a route
	// mux := http.NewServeMux()
	// mux.HandleFunc("/v1/healthcheck", app.healthcheckHandler)

	// Create our server
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,      // inactive connection
		ReadTimeout:  10 * time.Second, // time to read request body or header
		WriteTimeout: 10 * time.Second,
	}

	logger.Printf("Starting %s on port %d", cfg.env, cfg.port)
	err = srv.ListenAndServe()
	log.Fatal(err)
}

// Set up a database connection
func openDB(cfg config) (*sql.DB, error) {
	db, err := sql.Open("pgx", cfg.db.dsn)
	if err != nil {
		return nil, err
	}
	// Create a context
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	// Ping the database
	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}
	return db, err
}
