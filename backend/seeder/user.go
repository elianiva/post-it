package seeder

import (
	"context"
	"post-it-backend/business/models"
	"time"

	"github.com/jackc/pgx/v4"
)

var users = []models.User{
	{
		Username:  "manusia",
		Password:  "bernapas",
		FullName:  "Manusia Bernapas",
		AvatarURL: "https://avatars.dicebear.com/api/micah/manusia.svg",
		About:     "Nothing to see here...",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		Username:  "hooman",
		Password:  "notbreathing",
		FullName:  "Breathing Hooman",
		AvatarURL: "https://avatars.dicebear.com/api/micah/hooman.svg",
		About:     "Who am I..?",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		Username:  "guy420",
		Password:  "randomdude",
		FullName:  "Some Random Dude",
		AvatarURL: "https://avatars.dicebear.com/api/micah/dude.svg",
		About:     "Yo!",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
}

func SeedUser(ctx context.Context, conn *pgx.Conn) error {
	tx, err := conn.Begin(ctx)
	if err != nil {
		return err
	}

	// for _, u := range users {
	// 	u.Create()
	// }

	tx.Commit(ctx)

	return nil
}
