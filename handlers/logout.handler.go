package handlers

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

func LogoutHandler(ctx *fiber.Ctx) error {

	cookie := fiber.Cookie{
		Name: "jwt",
		Value: "",
		Expires: time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	ctx.Cookie(&cookie)

	return ctx.JSON(fiber.Map{
		"message" : "Success Logout",
	})

}