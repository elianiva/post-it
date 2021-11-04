package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	_ "post-it-backend/business/models"
	"post-it-backend/handlers"
	"post-it-backend/platform/migration"
	"post-it-backend/repository"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/jackc/pgx/v4"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	app := App()
	ctx := context.Background()

	config, err := pgx.ParseConfig(repository.DSN)
	if err != nil {
		log.Fatalf("Failed to parse config. Reason: %v", err)
	}

	conn, err := pgx.ConnectConfig(ctx, config)
	if err != nil {
		log.Fatalf("Failed to start connection. Reason: %v", err)
		os.Exit(1)
	}
	defer conn.Close(ctx)

	r := &handlers.Dependency{
		Conn: conn,
	}

	if os.Getenv("MIGRATE") == "true" {
		err = migration.Drop(ctx, r.Conn)
		if err != nil {
			log.Printf("Failed to drop table. Reason: %v", err)
			return
		}
		fmt.Println("======")
		err := migration.Migrate(ctx, r.Conn)
		if err != nil {
			log.Printf("Failed to migrate table. Reason: %v", err)
			return
		}
		return
	}

	registerRoutes(app, r)
	runApp(app)
}

func registerRoutes(app *fiber.App, r *handlers.Dependency) {
	// sanity check
	app.Get("/", r.Hello)

	api := app.Group("api", cors.New())
	api.Get("/users", r.AllUsers)
}

func runApp(app *fiber.App) {
	// catch SIGINT to shutdown gracefully
	go func() {
		s := make(chan os.Signal, 1)
		signal.Notify(s, os.Interrupt)
		<-s

		err := app.Shutdown()
		if err != nil {
			log.Fatalf("Failed to shutdown. Reason: %v", err)
		}
	}()

	err := app.Listen(":3000")
	if err != nil {
		log.Fatalf("Failed to start. Reason: %v", err)
	}
}
