package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ragilputri/absensi-pegawai/database"
	"github.com/ragilputri/absensi-pegawai/model"
)

func PegawaiHandlerGet(ctx *fiber.Ctx) error {
	userId := ctx.Locals("userId")
	var user model.User
	errUser := database.DB.First(&user, "id = ?", userId).Error
	if errUser != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "User Not Found",
			"error":   errUser,
		})
	}

	var role model.Role
	errRole := database.DB.First(&role, "id = ?", user.RoleRefer).Error
	if errRole != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "Role Not Found",
		})
	}

	responseRole := CreateResponseRole(role)
	responsePegawai := CreateResponseUser(user, responseRole)

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":   responsePegawai,
	})

}