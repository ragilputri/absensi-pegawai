package middlerware

import (
	"log"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/ragilputri/absensi-pegawai/database"
	"github.com/ragilputri/absensi-pegawai/model"
	"github.com/ragilputri/absensi-pegawai/handlers"
)

func AuthPegawai(ctx *fiber.Ctx) error {

	cookie := ctx.Cookies("jwt")

	token, errToken := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(handlers.SecretKey), nil
	})

	if errToken != nil{
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message":"Unauthenticated",
		})
	}

	claims := token.Claims.(*jwt.StandardClaims)

	var user model.User
	database.DB.First(&user, "id = ?", claims.Issuer)

	var role model.Role
	database.DB.First(&role, "id = ?", user.RoleRefer)
	log.Println(role)
	if role.Name != "pegawai"{
		return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message":"Forbidden Access",
		})
	}

	ctx.Locals("userId", claims.Issuer)

	return ctx.Next()

}

func AuthAdmin(ctx *fiber.Ctx) error {

	cookie := ctx.Cookies("jwt")

	token, errToken := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(handlers.SecretKey), nil
	})

	if errToken != nil{
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message":"Unauthenticated",
		})
	}

	claims := token.Claims.(*jwt.StandardClaims)

	var user model.User
	database.DB.First(&user, "id = ?", claims.Issuer)

	var role model.Role
	database.DB.First(&role, "id = ?", user.RoleRefer)
	log.Println(role)
	if role.Name != "admin"{
		return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message":"Forbidden Access",
		})
	}

	ctx.Locals("userId", claims.Issuer)

	return ctx.Next()

}