package main

import (
	"log"
	"os"
	"os/signal"
	"post-it-backend/handlers"
	"post-it-backend/prisma/db"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	app := App()

	// setup and connect prisma client to postgresql
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

	r := &handlers.Dependency{
		DB: client,
	}

	registerRoutes(r, app)
	runApp(app)
}

func registerRoutes(r *handlers.Dependency, app *fiber.App) {
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
