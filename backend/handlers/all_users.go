package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func (d *Dependency) AllUsers(c *fiber.Ctx) error {
	return c.SendString("HELLLOO")
}
