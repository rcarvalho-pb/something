package user_repository

import (
	"io"

	"github.com/jmoiron/sqlx"
)

type Database interface {
	io.Closer
	GetDB() *sqlx.DB
}
