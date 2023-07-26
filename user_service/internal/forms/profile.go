package form

type ProfileRequest struct {
	Firstname  string `json:"firstname" binding:"required"`
	Middlename string `json:"middlename"`
	Lastname   string `json:"lastname" binding:"required"`
}

type ChangePasswordRequest struct {
	OldPassword             string `json:"old_password" binding:"required"`
	NewPassword             string `json:"new_password" binding:"required"`
	NewPasswordConfirmation string `json:"new_password_confirmation" binding:"required"`
}
