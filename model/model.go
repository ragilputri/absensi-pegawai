package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Name      string         `json:"name"`
	Password  string         `json:"-" gorm:"column:password"`
	Email     string         `json:"email"`
	BirthDate  string         `json:"birth_date"`
	Address   string         `json:"address"`
	Phone     string         `json:"phone"`
	Photo	string			`json:"photo"`
	RoleRefer	string			`json:"role_refer"`
	Role		Role		`gorm:"foreignKey:RoleRefer"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time      `json:"update_at"`
	DeleteAt  gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

type Role struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Name      string         `json:"name"`
	CreatedAt time.Time        `json:"created_at"`
	UpdatedAt time.Time      `json:"update_at"`
	DeleteAt  gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

type Absen struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Date      string         `json:"date"`
	Masuk		string			`json:"masuk"`
	Keluar		string			`json:"keluar"`
	UserRefer	string			`json:"user_refer"`
	User		User		`gorm:"foreignKey:UserRefer"`
	CreatedAt time.Time        `json:"created_at"`
	UpdatedAt time.Time      `json:"update_at"`
	DeleteAt  gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

