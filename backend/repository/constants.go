package repository

import (
	"os"
	"strings"

	_ "github.com/joho/godotenv/autoload"
)

var DSN = strings.Join([]string{
	"user=" + os.Getenv("POSTGRES_USER"),
	"password=" + os.Getenv("POSTGRES_PASSWORD"),
	"host=" + os.Getenv("POSTGRES_HOST"),
	"port=" + os.Getenv("POSTGRES_PORT"),
	"dbname=" + os.Getenv("POSTGRES_DB"),
	"sslmode=disable",
}, " ")
