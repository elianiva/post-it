package models

import (
	"time"

	"github.com/aldy505/bob"
)

type Follow struct {
	ID          uint
	FollowingID string
	FollowerID  string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (_ Follow) GetSchema() (string, error) {
	sql, _, err := bob.
		CreateTable("follows").
		AddColumn(bob.ColumnDef{
			Name:   "id",
			Type:   "SERIAL",
			Extras: []string{"NOT NULL", "PRIMARY KEY"},
		}).
		AddColumn(bob.ColumnDef{
			Name:   "following_id",
			Type:   "TEXT",
			Extras: []string{"NOT NULL", `REFERENCES "users" ("username")`},
		}).
		AddColumn(bob.ColumnDef{
			Name:   "follower_id",
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

func (_ Follow) Drop() string {
	return `DROP TABLE IF EXISTS "follows" CASCADE`
}
