package handlers

import (
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/ragilputri/absensi-pegawai/database"
	"github.com/ragilputri/absensi-pegawai/model"
	"github.com/ragilputri/absensi-pegawai/model/request"
	"github.com/ragilputri/absensi-pegawai/utils"
)

const SecretKey = "secret"
func Loginhandler(ctx *fiber.Ctx) error {
	loginRequest := new(request.LoginRequest)
	if errLogin := ctx.BodyParser(loginRequest); errLogin != nil{
		return errLogin
	}

	validate := validator.New()
	errValidate := validate.Struct(loginRequest)
	if errValidate != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "failed",
			"error":   errValidate.Error(),
		})
	}

	var user model.User
	err := database.DB.First(&user, "email = ?", loginRequest.Email).Error
	if err != nil{
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message" : "Akun Tidak Terdaftar",
		})
	}

	isvalid := utils.CheckPasswordHash(loginRequest.Password, user.Password)
	if !isvalid{
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message" : "Password Salah",
		})
	}

	// claims := jwt.MapClaims{}
	// claims["id"] = user.ID
	// claims["name"] = user.Name
	// claims["email"] = user.Email
	// claims["role_refer"] = user.RoleRefer
	// claims["exp"] = time.Now().Add(time.Minute * 10).Unix()

	// token, errGenerated := utils.GenerateToken(&claims)
	// if errGenerated != nil{
	// 	return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
	// 		"message" : "Wrong Credential",
	// 	})
	// }

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer: strconv.Itoa(int(user.ID)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),

	})

	token, errToken := claims.SignedString([]byte(SecretKey))
	if errToken != nil{
		ctx.Status(fiber.StatusInternalServerError)
		return ctx.JSON(fiber.Map{
			"message":"cloud not login",
		})
	}

	cookie := fiber.Cookie{
		Name: "jwt",
		Value: token,
		Expires: time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	ctx.Cookie(&cookie)

	return ctx.JSON(fiber.Map{
		"Message" : "Success",
	})

}