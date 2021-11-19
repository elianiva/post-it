package migration

import (
	"context"
	"fmt"
	"post-it-backend/business/models"

	"github.com/jackc/pgx/v4"
)

var schemas = map[string]models.Model{
	// it needs to be in this specific order, otherwise it will fail
	"user":     models.User{},
	"follows":  models.Follow{},
	"post":     models.Post{},
	"comments": models.Comment{},
}

func Drop(ctx context.Context, conn *pgx.Conn) error {
	tx, err := conn.Begin(ctx)

	for name, entity := range schemas {
		fmt.Print("Dropping table " + name + "... ")
		drop, err := entity.Drop();
		if err != nil {
			return err;
		}

		_, err = tx.Exec(ctx, drop)
		if err != nil {
			return err
		}

		fmt.Println("Done!")
	}

	err = tx.Commit(ctx)
	if err != nil {
		return err
	}

	return nil
}

func Migrate(ctx context.Context, conn *pgx.Conn) error {
	tx, err := conn.Begin(ctx)
	if err != nil {
		return err
	}
	for name, entity := range schemas {
		fmt.Print("Migrating " + name + "... ")

		schema, err := entity.GetSchema()
		if err != nil {
			return err
		}

		_, err = tx.Exec(ctx, schema)
		if err != nil {
			return err
		}

		fmt.Println("Done!")
	}
	tx.Commit(ctx)

	return nil
}
