package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ragilputri/absensi-pegawai/handlers"
	"github.com/ragilputri/absensi-pegawai/middlerware"
	"github.com/ragilputri/absensi-pegawai/utils"
)

func RouteInit(r *fiber.App) {

	//Login Endpoint
	r.Post("/login", handlers.Loginhandler)

	//Roles Enpoint
	r.Get("/role", middlerware.AuthAdmin, handlers.RoleHandlerGet)
	r.Post("/role", middlerware.AuthAdmin, handlers.RoleHandlerCreate)
	r.Get("/role/:id",middlerware.AuthAdmin, handlers.RoleHandlerGetById)
	r.Put("/role/:id",middlerware.AuthAdmin, handlers.RoleHandlerUpdate)
	r.Delete("/role/:id",middlerware.AuthAdmin, handlers.RoleHandlerDelete)

	//Users Enpoint
	r.Get("/user", middlerware.AuthAdmin, handlers.UserHandlerGet)
	r.Post("/user", middlerware.AuthAdmin, utils.HandlerSingleFile, handlers.UserHandlerCreate)
	r.Get("/user/:id", middlerware.AuthAdmin, handlers.UserHandlerGetById)
	r.Put("/user/:id", middlerware.AuthAdmin, handlers.UserHandlerUpdate)
	r.Delete("/user/:id", middlerware.AuthAdmin, handlers.UserHandlerDelete)

	//Absen Enpoint
	r.Post("/absen", middlerware.AuthPegawai, handlers.AbsenMasukHandler)
	r.Put("/absen", middlerware.AuthPegawai, handlers.AbsenKeluarHandler)
	r.Get("/absen", middlerware.AuthPegawai, handlers.DaftarAbsenHandler)

	//Akun Pegawai Endpoint
	r.Get("/myaccount", middlerware.AuthPegawai, handlers.PegawaiHandlerGet)

	//Logout Endpoint
	r.Post("/logout", handlers.LogoutHandler)

	r.Get("/", handlers.UserHandlerGet)
}
