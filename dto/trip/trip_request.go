package tripdto

type CreateTripRequest struct {
	// Name     string `json:"name" form:"name" validate:"required"`
	// Email    string `json:"email" form:"email" validate:"required"`
	// Password string `json:"password" form:"password" validate:"required"`

	Title          string `json:"title" form:"title" validate:"required"`
	Country        int    `json:"country" form:"country" validate:"required"`
	Accomodation   string `json:"accomodation" form:"accomodation" validate:"required"`
	Transportation string `json:"transportation" form:"transportation" validate:"required"`
	Eat            string `json:"eat" form:"eat" validate:"required"`
	Day            int    `json:"day" form:"day" validate:"required"`
	Night          int    `json:"night" form:"night" validate:"required"`
	DateTrip       string `json:"date_trip" form:"date_trip" validate:"required"`
	Price          int    `json:"price" form:"price" validate:"required"`
	Quota          int    `json:"quota" form:"quota" validate:"required"`
	Description    string `json:"description" form:"description" validate:"required"`
	Image          string `json:"image" form:"image"`
}

// type UpdateUserRequest struct {
// 	Id       int    `json:"id" form:"id"`
// 	Name     string `json:"name" form:"name"`
// 	Email    string `json:"email" form:"email"`
// 	Password string `json:"password" form:"password"`
// }
