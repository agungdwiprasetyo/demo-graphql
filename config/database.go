package config

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type databaseConn struct{}

func (d *databaseConn) LoadReadDB() *sqlx.DB {
	descriptor := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("POSTGRES_READ_HOST"), os.Getenv("POSTGRES_READ_USER"), os.Getenv("POSTGRES_READ_PASSWORD"), os.Getenv("POSTGRES_READ_DB"))

	db, err := sqlx.Open("postgres", descriptor)
	if err != nil {
		log.Fatalf("\x1b[31;1mFailed load read database, %v\x1b[0m\n", err)
	}
	return db
}

func (d *databaseConn) LoadWriteDB() *sqlx.DB {
	descriptor := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("POSTGRES_WRITE_HOST"), os.Getenv("POSTGRES_WRITE_USER"), os.Getenv("POSTGRES_WRITE_PASSWORD"), os.Getenv("POSTGRES_WRITE_DB"))

	db, err := sqlx.Open("postgres", descriptor)
	if err != nil {
		log.Fatalf("\x1b[31;1mFailed load write database, %v\x1b[0m\n", err)
	}
	return db
}
