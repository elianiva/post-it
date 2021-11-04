package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"post-it-backend/prisma/db"
)

func main() {
	client := db.NewClient()
	err := client.Prisma.Connect()
	if err != nil {
		log.Fatalf("Prisma client failed to connect. Reason: %v", err)
		os.Exit(1)
	}

	defer func() {
		if err := client.Prisma.Disconnect(); err != nil {
			log.Fatalf("Prisma client failed to disconnect. Reason: %v", err)
		}
	}()

	ctx := context.Background()
	err = createUsers(client, ctx)
	if err != nil {
		log.Fatalf("Error seeding users. Reason: %v", err)
	}
}

func createUsers(client *db.PrismaClient, ctx context.Context) error {
	users := []map[string]interface{}{
		{
			"email":     db.User.Email.Set("manusia@bernapas.com"),
			"username":  db.User.Username.Set("manusia"),
			"password":  db.User.Password.Set("manusia1234"),
			"fullname":  db.User.FullName.Set("Manusia Bernapas"),
			"avatarUrl": db.User.AvatarURL.Set("https://avatars.dicebear.com/api/micah/manusia.svg"),
			"about":     db.User.About.Set("Your friendly neighbourhood breathing hooman"),
		},
		{
			"email":     db.User.Email.Set("foo@bar.com"),
			"username":  db.User.Username.Set("FooBar"),
			"password":  db.User.Password.Set("foobar1234"),
			"fullname":  db.User.FullName.Set("Foo Bar Baz"),
			"avatarUrl": db.User.AvatarURL.Set("https://avatars.dicebear.com/api/micah/FooBar.svg"),
			"about":     db.User.About.Set("You thought it was foo bar, but it was me! Asdf"),
		},
		{
			"email":     db.User.Email.Set("hooman@business.co"),
			"username":  db.User.Username.Set("hooman"),
			"password":  db.User.Password.Set("hooman1234"),
			"fullname":  db.User.FullName.Set("Some Random Hooman"),
			"avatarUrl": db.User.AvatarURL.Set("https://avatars.dicebear.com/api/micah/FooBar.svg"),
		},
	}

	for _, u := range users {
		post, err := client.User.CreateOne(
			u["email"].(db.UserWithPrismaEmailSetParam),
			u["username"].(db.UserWithPrismaUsernameSetParam),
			u["password"].(db.UserWithPrismaPasswordSetParam),
			u["fullname"].(db.UserWithPrismaFullNameSetParam),
			u["avatarUrl"].(db.UserWithPrismaAvatarURLSetParam),
			u["about"].(db.UserWithPrismaAboutSetParam),
		).Exec(ctx)
		if err != nil {
			return err
		}

		result, _ := json.MarshalIndent(post, "", "  ")
		fmt.Printf("post: %s\n", result)
	}

	return nil
}
