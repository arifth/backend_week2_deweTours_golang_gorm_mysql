package countrydto

type CreateCountryRequest struct {
	Name string `json:"name" form:"name" validate:"required"`
}

// type UpdateCountryRequest struct {
// 	Id       int    `json:"id" form:"id"`
// 	Name     string `json:"name" form:"name"`
// 	Email    string `json:"email" form:"email"`
// 	Password string `json:"password" form:"password"`
// }
