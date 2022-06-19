package handlers

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/ragilputri/absensi-pegawai/database"
	"github.com/ragilputri/absensi-pegawai/model"
	"github.com/ragilputri/absensi-pegawai/model/request"
	"github.com/ragilputri/absensi-pegawai/model/response"
	"github.com/ragilputri/absensi-pegawai/utils"
)

func CreateResponseUser(userModel model.User, roleResponse response.RoleResponse) response.UserResponse {
	return response.UserResponse{
		ID:        userModel.ID,
		Name:      userModel.Name,
		Password:  userModel.Password,
		Email:     userModel.Email,
		BirthDate: userModel.BirthDate,
		Address:   userModel.Address,
		Phone:     userModel.Phone,
		Photo:     userModel.Photo,
		Role:      roleResponse,
	}
}

func UserHandlerGet(ctx *fiber.Ctx) error {
	var users []model.User
	database.DB.Find(&users)

	responseUsers := []response.UserResponse{}

	for _, user := range users {
		var role model.Role
		database.DB.First(&role, "id = ?", user.RoleRefer)
		responseUser := CreateResponseUser(user, CreateResponseRole(role))
		responseUsers = append(responseUsers, responseUser)
	}

	return ctx.JSON(fiber.Map{
		"data": responseUsers,
	})
}

func UserHandlerCreate(ctx *fiber.Ctx) error {
	user := new(request.UserCreateRequest)
	if err := ctx.BodyParser(user); err != nil {
		return err
	}

	//Validate User
	validate := validator.New()
	errValidate := validate.Struct(user)
	if errValidate != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "failed",
			"error":   errValidate.Error(),
		})
	}

	//image
	var filenameString string
	filename := ctx.Locals("filename")
	if filename == nil {
		return ctx.Status(422).JSON(fiber.Map{
			"message": "Photo is required",
		})
	} else {

		filenameString = fmt.Sprintf("%v", filename)
	}

	//Role
	var role model.Role
	errRole := database.DB.First(&role, "id = ?", user.RoleRefer).Error
	if errRole != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "Role Not Found",
		})
	}

	//Save User
	newUser := model.User{
		Name:      user.Name,
		Email:     user.Email,
		BirthDate: user.BirthDate,
		Address:   user.Address,
		Phone:     user.Phone,
		Photo:     filenameString,
		RoleRefer: user.RoleRefer,
	}

	hashedPassword, errPass := utils.HashingPassword(user.Password)
	if errPass != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "internal server error",
			"error":   errPass,
		})
	}
	newUser.Password = hashedPassword

	errCreatedUser := database.DB.Create(&newUser).Error
	if errCreatedUser != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "failed to store data",
		})
	}

	responseRole := CreateResponseRole(role)
	responseUser := CreateResponseUser(newUser, responseRole)

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    responseUser,
	})
}

func UserHandlerGetById(ctx *fiber.Ctx) error {
	userId := ctx.Params("id")
	var user model.User

	err := database.DB.First(&user, "id = ?", userId).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "User Not Found",
			"error":   err,
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
	responseUser := CreateResponseUser(user, responseRole)

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":   responseUser,
	})
}

func UserHandlerUpdate(ctx *fiber.Ctx) error {
	userRequest := new(request.UserUpdateRequest)
	err := ctx.BodyParser(userRequest)
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "Bad Request",
		})
	}

	var user model.User
	userId := ctx.Params("id")
	errFind := database.DB.First(&user, "id = ?", userId).Error
	if errFind != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "User Not Found",
		})
	}

	var role model.Role
	errRole := database.DB.First(&role, "id = ?", userRequest.RoleRefer).Error
	if errRole != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "Role Not Found",
		})
	}

	if userRequest.Photo != "" {
		user.Photo = userRequest.Photo
	}

	user.Name = userRequest.Name
	user.Email = userRequest.Email
	user.BirthDate = userRequest.BirthDate
	user.Address = userRequest.Address
	user.Phone = userRequest.Phone
	user.RoleRefer = userRequest.RoleRefer

	errUpdate := database.DB.Save(&user).Error
	if errUpdate != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "Internal Server Error",
		})
	}

	responseUser := CreateResponseUser(user, CreateResponseRole(role))

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    responseUser,
	})
}

func UserHandlerDelete(ctx *fiber.Ctx) error {
	userId := ctx.Params("id")
	var user model.User
	err := database.DB.First(&user, "id = ?", userId).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "failed",
			"error":   err.Error(),
		})
	}

	errDelete := database.DB.Delete(&user).Error
	if errDelete != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "Internal Server Error",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "User Was Deleted",
	})
}
