package handlers

import "github.com/gofiber/fiber/v2"

func (d *Dependency) Hello(ctx *fiber.Ctx) error {
	return ctx.SendString("Hello, 世界！")
}
