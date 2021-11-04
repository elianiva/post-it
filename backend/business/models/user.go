package models

import (
	"time"

	"github.com/aldy505/bob"
)

type User struct {
	Username  string
	Password  string
	FullName  string
	AvatarURL string `db:"avatar_url"`
	About     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (_ User) GetSchema() (string, error) {
	sql, _, err := bob.
		CreateTable("users").
		TextColumn("username", "NOT NULL", "UNIQUE", "PRIMARY KEY").
		TextColumn("password", "NOT NULL").
		TextColumn("fullname", "NOT NULL").
		TextColumn("avatar_url", "NOT NULL").
		TimeStampColumn("created_at", "NOT NULL", "DEFAULT CURRENT_TIMESTAMP").
		TimeStampColumn("updated_at", "NOT NULL").
		ToSql()
	if err != nil {
		return "", err
	}

	return sql, nil
}

func (_ User) Drop() string {
	return `DROP TABLE IF EXISTS "users" CASCADE`
}
