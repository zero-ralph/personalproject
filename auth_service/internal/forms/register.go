package form

type ProfileRequest struct {
	Firstname  string `json:"firstname" binding:"required"`
	Middlename string `json:"middlename"`
	Lastname   string `json:"lastname" binding:"required"`
}

type RegisterRequest struct {
	Username             string `json:"username" binding:"required"`
	Password             string `json:"password" binding:"required,min=8"`
	PasswordConfirmation string `json:"password_confirmation" binding:"required"`
	Profile              ProfileRequest
}
