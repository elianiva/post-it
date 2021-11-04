package models

import (
	"time"

	"github.com/aldy505/bob"
)

type Post struct {
	ID        uint
	Content   string
	UserID    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (_ Post) GetSchema() (string, error) {
	sql, _, err := bob.
		CreateTableIfNotExists("posts").
		AddColumn(bob.ColumnDef{
			Name:   "id",
			Type:   "SERIAL",
			Extras: []string{"NOT NULL", "PRIMARY KEY"},
		}).
		TextColumn("content").
		AddColumn(bob.ColumnDef{
			Name:   "user_id",
			Type:   "TEXT",
			Extras: []string{"NOT NULL", `REFERENCES "users" ("username")`},
		}).
		TimeStampColumn("created_at", "NOT NULL", "DEFAULT CURRENT_TIMESTAMP").
		TimeStampColumn("updated_at", "NOT NULL").
		ToSql()
	if err != nil {
		return "", err
	}

	return sql, nil
}

func (_ Post) Drop() string {
	return `DROP TABLE IF EXISTS "users" CASCADE`
}
