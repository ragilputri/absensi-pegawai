package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/ragilputri/absensi-pegawai/database"
	"github.com/ragilputri/absensi-pegawai/migration"
	"github.com/ragilputri/absensi-pegawai/route"
)

func main() {
	database.DatabaseInit()

	migration.RunMigration()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	route.RouteInit(app)

	app.Listen(":3000")
}