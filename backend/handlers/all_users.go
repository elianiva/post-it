package handlers

import (
	"log"
	"net/http"
	"post-it-backend/prisma/db"

	"github.com/gofiber/fiber/v2"
	"github.com/thoas/go-funk"
)

func (d *Dependency) AllUsers(c *fiber.Ctx) error {
	u, err := d.DB.User.FindMany().Exec(c.Context())
	if err != nil {
		log.Printf("Failed to execute users query. Reason: %v", err)
		return c.Status(http.StatusInternalServerError).Send([]byte("Internal Server Error: Failed to execute db query."))
	}

	users := funk.Map(u, func(u db.UserModel) map[string]string {
		about, _ := u.About()

		return map[string]string{
			"email":      u.Email,
			"username":   u.Username,
			"full_name":  u.FullName,
			"avatar_url": u.AvatarURL,
			"about":      about,
		}
	})

	return c.JSON(map[string]interface{}{
		"status": http.StatusOK,
		"data":   users,
	})
}
