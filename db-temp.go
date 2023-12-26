package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

func connectTemp() (*sql.DB, error) {
	// Capture connection properties.
	cfg := mysql.Config{
		User:   os.Getenv("DO_AL_CORE_DB_USER"),
		Passwd: os.Getenv("DO_AL_CORE_DB_PASS"),
		Net:    "tcp",
		Addr:   os.Getenv("DO_AL_CORE_DB_HOST") + ":" + os.Getenv("DO_AL_CORE_DB_PORT"),
		DBName: os.Getenv("DO_AL_CORE_DB_NAME"),
	}

	// Get a database handle.
	var err error

	db, err = sql.Open("mysql", cfg.FormatDSN())

	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()

	if pingErr != nil {
		log.Fatal(pingErr)
	}

	fmt.Println("Connected!")

	return db, nil
}
