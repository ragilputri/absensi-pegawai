package database

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
)

var DB *gorm.DB
var err error

func DatabaseInit() {
	dsn := "host=localhost user=postgres password=ragil312 dbname=absensi_pegawai port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Cannot connect to database")
	}
	fmt.Println("Connected to database")
}
