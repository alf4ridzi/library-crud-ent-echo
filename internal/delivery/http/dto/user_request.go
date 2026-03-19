package dto

type UserChangePasswordRequest struct {
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

type UserUpdateRequest struct {
	Name     *string `json:"name"`
	Email    *string `json:"email"`
	Username *string `json:"username"`
}
