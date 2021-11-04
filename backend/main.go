package main

import (
	"log"
	"os"
	"os/signal"
	h "post-it-backend/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	app := App()

	registerRoutes(app)
	runApp(app)
}

func registerRoutes(app *fiber.App) {
	app.Use(cors.New())

	app.Get("/", h.Hello)
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
