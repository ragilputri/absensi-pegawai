package handlers

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/ragilputri/absensi-pegawai/database"
	"github.com/ragilputri/absensi-pegawai/model"
	"github.com/ragilputri/absensi-pegawai/model/request"
	"github.com/ragilputri/absensi-pegawai/model/response"
)

func CreateResponseRole(roleModel model.Role) response.RoleResponse  {
	return response.RoleResponse{
		ID: roleModel.ID,
		Name : roleModel.Name,
	}
}

func RoleHandlerGet(ctx *fiber.Ctx)error  {
	var roles []model.Role
	err := database.DB.Find(&roles).Error

	if err != nil{
		return ctx.Status(400).JSON(fiber.Map{
			"message": "Failed",
			"error":   err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"data":   roles,
	})
}

func RoleHandlerCreate(ctx *fiber.Ctx) error {
	role := new(request.RoleCreateRequest)
	if err := ctx.BodyParser(role); err != nil {
		return err
	}

	//Validate Role
	validate := validator.New()
	errValidate := validate.Struct(role)
	if errValidate != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "failed",
			"error":   errValidate.Error(),
		})
	}

	//Save Role
	newRole := model.Role{
		Name:    role.Name,
	}

	errCreatedRole := database.DB.Create(&newRole).Error
	if errCreatedRole != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "failed to store data",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    newRole,
	})
}

func RoleHandlerGetById(ctx *fiber.Ctx)error  {
	roleId := ctx.Params("id")
	var role model.Role

	err := database.DB.First(&role, "id = ?", roleId).Error
	if err != nil{
		return ctx.Status(404).JSON(fiber.Map{
			"message": "Data Not Found",
			"error":   err,
		})
	}
	
	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":   role,
	})
}

func RoleHandlerUpdate(ctx *fiber.Ctx)error  {
	roleRequest := new(request.RoleCreateRequest)
	err := ctx.BodyParser(roleRequest)
	if err != nil{
		return ctx.Status(400).JSON(fiber.Map{
			"message": "Bad Request",
		})
	}

	var role model.Role
	roleId := ctx.Params("id")
	errFind := database.DB.First(&role, "id = ?", roleId).Error
	if errFind != nil{
		return ctx.Status(404).JSON(fiber.Map{
			"message": "Data Not Found",
		})
	}

	role.Name = roleRequest.Name

	errUpdate := database.DB.Save(&role).Error
	if errUpdate != nil{
		return ctx.Status(500).JSON(fiber.Map{
			"message": "Internal Server Error",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data"	: role,
	})
}

func RoleHandlerDelete(ctx *fiber.Ctx)error  {
	roleId := ctx.Params("id")
	var role model.Role
	err := database.DB.First(&role, "id = ?", roleId).Error
	if err != nil{
		return ctx.Status(404).JSON(fiber.Map{
			"message": "failed",
			"error":   err.Error(),
		})
	}

	errDelete := database.DB.Delete(&role).Error
	if errDelete != nil{
		return ctx.Status(500).JSON(fiber.Map{
			"message": "Internal Server Error",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "Role Was Deleted",
	})
}