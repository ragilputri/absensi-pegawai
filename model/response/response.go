package response


type UserResponse struct {
	ID        uint         `json:"id"`
	Name      string       `json:"name"`
	Password  string       `json:"-" gorm:"column:password"`
	Email     string       `json:"email"`
	BirthDate string        `json:"birth_date"`
	Address   string       `json:"address"`
	Phone     string       `json:"phone"`
	Photo     string       `json:"photo"`
	Role      RoleResponse `json:"role_refer"`
}

type RoleResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type AbsenResponse struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Date      string         `json:"date"`
	Masuk		string			`json:"masuk"`
	Keluar		string			`json:"keluar"`
	User		UserResponse		`json:"user_refer"`
}

type PersonAbsenResponse struct {
	Date      string         `json:"date"`
	Masuk		string			`json:"masuk"`
	Keluar		string			`json:"keluar"`
}
