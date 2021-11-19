package models

import (
	"time"

	"github.com/aldy505/bob"
)

type Reaction struct {
	ID        uint
	PostID    uint
	UserID    string
	Type      int
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (_ Reaction) GetSchema() (string, error) {
	sql, _, err := bob.
		CreateTable("reactions").
		AddColumn(bob.ColumnDef{
			Name:   "id",
			Type:   "SERIAL",
			Extras: []string{"NOT NULL", "PRIMARY KEY"},
		}).
		AddColumn(bob.ColumnDef{
			Name:   "post_id",
			Type:   "INT",
			Extras: []string{"NOT NULL", `REFERENCES "posts" ("id")`},
		}).
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

func (_ Reaction) Drop() (string, error) {
	sql, _, err := bob.DropTableIfExists("reactions").Cascade().ToSql()
	if err != nil {
		return "", err
	}

	return sql, nil
}
