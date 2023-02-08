//Filename: cmd/api/main.go

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

// setup a struct to hold server configueration
type config struct {
	port int
	env  string
}

// Setup dependency injection
type application struct {
	config config
	logger *log.Logger
}

//setup main() function

func main() {
	var cfg config

	//Get the arguments from the user for the
	//server configuration

	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment(development|staging|production)")
	flag.Parse()

	//create a logger

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	//create an object of type application

	app := &application{
		config: cfg,
		logger: logger,
	}

	//create our server

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,      //inactive connection
		ReadTimeout:  10 * time.Second, //time to read request body or header
		WriteTimeout: 10 * time.Second,
	}

	//start our server

	logger.Printf("starting %s server on port %d", cfg.env, cfg.port)
	err := srv.ListenAndServe()
	log.Fatal(err)
}
