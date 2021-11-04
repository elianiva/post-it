package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"post-it-backend/prisma/db"
)

func main() {
	client := db.NewClient()
	err := client.Prisma.Connect()
	if err != nil {
		log.Fatalf("Prisma client failed to connect. Reason: %v", err)
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
			"avatarUrl": nil,
		},
		{
			"email":     db.User.Email.Set("foo@bar.com"),
			"username":  db.User.Username.Set("FooBar"),
			"password":  db.User.Password.Set("foobar1234"),
			"fullname":  db.User.FullName.Set("Foo Bar Baz"),
			"avatarUrl": nil,
		},
		{
			"email":     db.User.Email.Set("foo@bar.com"),
			"username":  db.User.Username.Set("FooBar"),
			"password":  db.User.Password.Set("foobar1234"),
			"fullname":  db.User.FullName.Set("Foo Bar Baz"),
			"avatarUrl": nil,
		},
	}

	for _, u := range users {
		post, err := client.User.CreateOne(
			u["email"].(db.UserWithPrismaEmailSetParam),
			u["username"].(db.UserWithPrismaUsernameSetParam),
			u["password"].(db.UserWithPrismaPasswordSetParam),
			u["fullname"].(db.UserWithPrismaFullNameSetParam),
			u["avatarUrl"].(db.UserWithPrismaAvatarURLSetParam),
		).Exec(ctx)
		if err != nil {
			return err
		}

		result, _ := json.MarshalIndent(post, "", "  ")
		fmt.Printf("post: %s\n", result)
	}

	return nil
}
