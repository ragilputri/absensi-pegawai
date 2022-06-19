package utils

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

const DefaultPathAssetImage = "./public/photo/"

func HandlerSingleFile(ctx *fiber.Ctx) error {
	file, errFile := ctx.FormFile("photo")
	if errFile != nil {
		log.Println("Error File : ", errFile)
	}

	var filename *string

	if file != nil {
		filename = &file.Filename

		errSaveFile := ctx.SaveFile(file, fmt.Sprint("./public/photo/", *filename))
		if errSaveFile != nil {
			return ctx.Status(400).JSON(fiber.Map{
				"message": "failed save file",
				"error":   errSaveFile.Error(),
			})
		}
	} else {
		log.Println("Nothing File To Uploading")
	}

	if filename != nil {
		ctx.Locals("filename", *filename)
	} else {
		ctx.Locals("filename", nil)
	}
	return ctx.Next()
}