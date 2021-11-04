package models

import (
	"time"

	"github.com/aldy505/bob"
)

type Comment struct {
	ID        uint
	PostID    uint
	UserID    string
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (_ Comment) GetSchema() (string, error) {
	sql, _, err := bob.
		CreateTable("comments").
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
		TextColumn("content", "NOT NULL").
		TimeStampColumn("created_at", "NOT NULL", "DEFAULT CURRENT_TIMESTAMP").
		TimeStampColumn("updated_at", "NOT NULL").
		ToSql()
	if err != nil {
		return "", err
	}

	return sql, nil
}

func (_ Comment) Drop() string {
	return `DROP TABLE IF EXISTS "comments" CASCADE`
}
