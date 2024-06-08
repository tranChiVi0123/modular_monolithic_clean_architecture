package presenter

type UserLogin struct {
	EmailAddress string `json:"email"`
	Password     string `json:"password"`
}

type UserResponse struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	BirthDate    string `json:"birth_date"`
	EmailAddress string `json:"email_address"`
	Address      string `json:"address"`
}

type UserRegister struct {
	UserName     string `json:"username"`
	EmailAddress string `json:"email"`
	Password     string `json:"password"`
}
