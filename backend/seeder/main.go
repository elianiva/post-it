package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"post-it-backend/prisma/db"
)

var users = []map[string]string{
	{
		"email":     "manusia@bernapas.com",
		"username":  "manusia",
		"password":  "manusia1234",
		"fullname":  "Manusia Bernapas",
		"avatarUrl": "https://avatars.dicebear.com/api/micah/manusia.svg",
		"about":     "Your friendly neighbourhood breathing hooman",
	},
	{
		"email":     "foo@bar.com",
		"username":  "FooBar",
		"password":  "foobar1234",
		"fullname":  "Foo Bar Baz",
		"avatarUrl": "https://avatars.dicebear.com/api/micah/FooBar.svg",
		"about":     "You thought it was foo bar, but it was me! Asdf",
	},
	{
		"email":     "hooman@business.co",
		"username":  "hooman",
		"password":  "hooman1234",
		"fullname":  "Some Random Hooman",
		"avatarUrl": "https://avatars.dicebear.com/api/micah/FooBar.svg",
		"about":     "",
	},
}

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
	err = userSeeder(client, ctx)
	if err != nil {
		log.Fatalf("Error seeding users. Reason: %v", err)
	}
}

func userSeeder(client *db.PrismaClient, ctx context.Context) error {
	for _, u := range users {
		err := createUser(ctx, client, u)
		if err != nil {
			log.Printf("Failed to seed user. Reason: %v", err)
		}
	}

	return nil
}

func createUser(ctx context.Context, client *db.PrismaClient, data map[string]string) error {
	var post *db.UserModel
	var err error

	if data["about"] != "" {
		post, err = client.User.CreateOne(
			// required fields
			db.User.Email.Set(data["email"]),
			db.User.Username.Set(data["username"]),
			db.User.Password.Set(data["password"]),
			db.User.FullName.Set(data["fullname"]),
			db.User.AvatarURL.Set(data["avatarUrl"]),

			// optional fields
			db.User.About.Set(data["about"]),
		).Exec(ctx)
	} else {
		post, err = client.User.CreateOne(
			// required fields
			db.User.Email.Set(data["email"]),
			db.User.Username.Set(data["username"]),
			db.User.Password.Set(data["password"]),
			db.User.FullName.Set(data["fullname"]),
			db.User.AvatarURL.Set(data["avatarUrl"]),
		).Exec(ctx)
	}

	if err != nil {
		return err
	}

	result, _ := json.MarshalIndent(post, "", "  ")
	fmt.Printf("post: %s\n", result)

	return nil
}
