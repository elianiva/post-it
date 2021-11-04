package main

import (
	"github.com/gofiber/fiber/v2"
)

func App() *fiber.App {
	app := fiber.New(fiber.Config{
		AppName:       "Post-It",
		CaseSensitive: true,
	})

	return app
}
