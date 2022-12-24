package usersdto

type UserResponse struct {
	ID       uint   `json:"id"`
	Name     string `json:"name" form:"name" validate:"required"`
	Email    string `json:"email" form:"email" validate"required"`
	Password string `json:"password" form:"password" validate:"required"`
}
