package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

const version = "0.0.1"
const defaultPort = 8081

type config struct {
	port int
	env  string
}

type application struct {
	config config
	logger *log.Logger
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", defaultPort, "API Server Port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	app := &application{
		config: cfg,
		logger: logger,
	}

	const readTimeout = 10 * time.Second

	mux := http.NewServeMux()
	mux.HandleFunc("/v1/healthcheck", app.healthcheckHandler)

	srv := &http.Server{
		Addr:              fmt.Sprintf(":%d", cfg.port),
		Handler:           mux,
		IdleTimeout:       time.Minute,
		ReadTimeout:       readTimeout,
		WriteTimeout:      time.Second,
		ReadHeaderTimeout: time.Second,
	}

	logger.Printf("starting %s server on %d", cfg.env, cfg.port)

	err := srv.ListenAndServe()
	logger.Fatal(err)
}
