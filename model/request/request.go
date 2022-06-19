package request

type UserCreateRequest struct {
	Name      string `json:"name"`
	Password  string `json:"-" gorm:"column:password"`
	Email     string `json:"email"`
	BirthDate string `json:"birth_date" form:"birth_date"`
	Address   string `json:"address"`
	Phone     string `json:"phone"`
	Photo     string `json:"photo"`
	RoleRefer string `json:"role_refer" form:"role_refer"`
}

type UserUpdateRequest struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	BirthDate string `json:"birth_date" form:"birth_date"`
	Address   string `json:"address"`
	Phone     string `json:"phone"`
	Photo     string `json:"photo"`
	RoleRefer string `json:"role_refer" form:"role_refer"`
}

type AbsenCreateRequest struct {
	Date      string      `json:"date"`
	Masuk     string    `json:"masuk"`
	Keluar    string     `json:"keluar"`
	UserRefer string         `json:"user_refer"`
}

type UserChangePasswordRequest struct {
	Password string `json:"-" gorm:"column:password"`
}

type RoleCreateRequest struct {
	Name string `json:"name" validate:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate :"required"`
}

