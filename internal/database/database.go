package database

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type Database interface {
	GetDB() *pgxpool.Pool
	CloseDB()
}
