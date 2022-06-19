package handlers

import (
	"log"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/ragilputri/absensi-pegawai/database"
	"github.com/ragilputri/absensi-pegawai/model"
	"github.com/ragilputri/absensi-pegawai/model/response"
)

func CreateResponseAbsen(absenModel model.Absen, userResponse response.UserResponse) response.AbsenResponse {
	return response.AbsenResponse{
		ID:     absenModel.ID,
		Date:   absenModel.Date,
		Masuk:  absenModel.Masuk,
		Keluar: absenModel.Keluar,
		User:   userResponse,
	}
}

func CreateResponseAbsenPerson(absenModel model.Absen) response.PersonAbsenResponse {
	return response.PersonAbsenResponse{
		Date:   absenModel.Date,
		Masuk:  absenModel.Masuk,
		Keluar: absenModel.Keluar,
	}
}

func AbsenMasukHandler(ctx *fiber.Ctx) error {
	userId := ctx.Locals("userId")
	log.Println("userId", userId)
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
			"error":   errRole,
		})
	}

	var now = time.Now()
	var absen model.Absen
	errAbsen := database.DB.First(&absen, "user_refer = ? AND date = ?", user.ID, now.Format("2006-01-02")).Error
	log.Println("errAbsen", errAbsen)
	log.Println("Absen", absen)
	if errAbsen != nil {
		//Save Absen Masuk
		absenMasuk := model.Absen{
			Date:      now.Format("2006-01-02"),
			Masuk:     now.Format("15:04:05"),
			UserRefer: strconv.Itoa(int(user.ID)),
		}
		errCreatedAbsen := database.DB.Create(&absenMasuk).Error
		if errCreatedAbsen != nil {
			return ctx.Status(500).JSON(fiber.Map{
				"message": "failed to store data",
			})
		}

		responseAbsen := CreateResponseAbsen(absenMasuk, CreateResponseUser(user, CreateResponseRole(role)))

		return ctx.JSON(fiber.Map{
			"message": "success",
			"data":    responseAbsen,
		})

	}

	return ctx.JSON(fiber.Map{
		"message": "Anda Sudah Absen Masuk",
	})

}

func AbsenKeluarHandler(ctx *fiber.Ctx) error {
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
			"error":   errRole,
		})
	}

	var now = time.Now()
	var absen model.Absen
	errAbsen := database.DB.First(&absen, "user_refer = ? AND date = ?", user.ID, now.Format("2006-01-02")).Error
	if errAbsen != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "Anda Belum Absen Masuk",
			"error":   errAbsen,
		})
	}

	if absen.Keluar != "" {
		return ctx.JSON(fiber.Map{
			"message": "Anda Sudah Absen Pulang",
		})
	} else if absen.Keluar == "" {
		absen.Keluar = now.Format("15:04:05")
	}

	errUpdateAbsen := database.DB.Save(&absen).Error
	if errUpdateAbsen != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "failed to store data",
		})
	}

	responseAbsen := CreateResponseAbsen(absen, CreateResponseUser(user, CreateResponseRole(role)))

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    responseAbsen,
	})

}

func DaftarAbsenHandler(ctx *fiber.Ctx) error {
	userId := ctx.Locals("userId")
	var user model.User
	errUser := database.DB.First(&user, "id = ?", userId).Error
	if errUser != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "User Not Found",
			"error":   errUser,
		})
	}

	var absens []model.Absen
	errAbsen := database.DB.Select("date", "masuk", "keluar").Where("user_refer = ?", user.ID).Find(&absens).Error
	if errAbsen != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "Absens Not Found",
			"error":   errAbsen,
		})
	}

	responseAbsenPerson := []response.PersonAbsenResponse{}

	for _, absen := range absens {
		responsePerson := CreateResponseAbsenPerson(absen)
		responseAbsenPerson = append(responseAbsenPerson, responsePerson)
	}

	return ctx.JSON(fiber.Map{
		"data": responseAbsenPerson,
	})

}
