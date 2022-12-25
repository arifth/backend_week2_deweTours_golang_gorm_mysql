package tripdto

type CreateTripRequest struct {
	// Name     string `json:"name" form:"name" validate:"required"`
	// Email    string `json:"email" form:"email" validate:"required"`
	// Password string `json:"password" form:"password" validate:"required"`

	Title          string `json:"title" form:"title" `
	Country        int    `json:"country" form:"country"`
	Accomodation   string `json:"accomodation" form:"accomodation" `
	Transportation string `json:"transportation" form:"transportation" `
	Eat            string `json:"eat" form:"eat" `
	Day            int    `json:"day" form:"day" `
	Night          int    `json:"night" form:"night" `
	DateTrip       string `json:"date_trip" form:"date_trip" `
	Price          int    `json:"price" form:"price" `
	Quota          int    `json:"quota" form:"quota" `
	Description    string `json:"description" form:"description" `
	Image          string `json:"image" form:"image"`
}

// type UpdateUserRequest struct {
// 	Id       int    `json:"id" form:"id"`
// 	Name     string `json:"name" form:"name"`
// 	Email    string `json:"email" form:"email"`
// 	Password string `json:"password" form:"password"`
// }
