package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"reading_list/internal/data"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

type config struct {
	port int
	env  string
	dsn  string
}

type application struct {
	config config
	logger *log.Logger
	models data.Models
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "API Server Port")
	flag.StringVar(&cfg.env, "env", "dev", "Environment (dev|stage|prod)")
	flag.StringVar(&cfg.dsn, "db-dsn", os.Getenv("READINGLIST_DB_DSN"), "PostgreSQL DSN")
	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	db, err := sql.Open("postgres", cfg.dsn)
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
		return
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatalf("Could not ping the database: %v", err)
		return
	}

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
		models: data.NewModels(db),
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
