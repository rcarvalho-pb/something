package sqlite3db

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func NewSqlite3db(path string) *sqlx.DB {
	db, err := sqlx.Connect("sqlite3", path)
	if err != nil {
		log.Fatal("error connecting to database")
	}

	return db
}
