package migration

import (
	"fmt"
	"log"

	"github.com/ragilputri/absensi-pegawai/database"
	"github.com/ragilputri/absensi-pegawai/model"
)

func RunMigration() {
	err := database.DB.AutoMigrate(&model.User{}, &model.Role{}, &model.Absen{})
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Database Migrated")
}