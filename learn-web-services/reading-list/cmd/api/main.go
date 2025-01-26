package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

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

	flag.IntVar(&cfg.port, "port", 4000, "API Server Port")
	flag.StringVar(&cfg.env, "env", "dev", "Environment (dev|stage|prod)")
	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	dbUrl := "postgres://postgres:superSecretPassword!@localhost:5432/reading_list?sslmode=disable"
	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
		return
	}
	defer db.Close()

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatalf("Could not create database driver: %v", err)
		return
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://../../db/migrations", // Path to migration files
		"postgres",                   // Database name
		driver,
	)
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
		return
	}

	err = m.Up()
	if err != nil {
		if err == migrate.ErrNoChange {
			log.Print("No DB Migrations")
		} else {
			log.Fatalf("Could not apply migrations: %v", err)
			return
		}
	} else {
		log.Print("DB Migrations Applied")
	}

	app := &application{
		config: cfg,
		logger: logger,
	}
	addr := fmt.Sprintf(":%d", app.config.port)
	srv := &http.Server{
		Addr:         addr,
		Handler:      app.route(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Printf("Starting %s server on %d", app.config.env, app.config.port)
	err = srv.ListenAndServe()
	logger.Fatal(err)
}
