package handlers

import (
	"github.com/jackc/pgx/v4"
)

type Dependency struct {
	Conn *pgx.Conn
}
