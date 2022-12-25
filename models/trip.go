package models

// type Trip struct {
// 	gorm.Model
// 	Title          string  `json:"title"`
// 	CountryId      int     `json:"country_id"`
// 	Country        Country `json:"country" gorm:"constraint:OnUpdate:CASCADE,one2one:country"`
// 	Accomodation   string  `json:"accomodation"`
// 	Transportation string  `json:"transportation"`
// 	Eat            string  `json:"eat"`
// 	Day            int     `json:"day"`
// 	Night          int     `json:"night"`
// 	DateTrip       string  `json:"date_trip"`
// 	Price          int     `json:"price"`
// 	Quota          int     `json:"quota"`
// 	Description    string  `json:"description"`
// 	Image          string  `json:"image"`
// }

type Trip struct {
	ID             int             `json:"id"  gorm:"primary_key:auto_increment" `
	Title          string          `json:"title" gorm:"type : varchar (255)"`
	Country_id     int             `json:"-"`
	Country        CountryResponse `json:"country"`
	Accomodation   string          `json:"accomodation" gorm:"type : varchar (255)"`
	Transportation string          `json:"transportation" gorm:"type : varchar (255)"`
	Eat            string          `json:"eat" gorm:"type : varchar (255)"`
	Day            int             `json:"day" gorm:"type : varchar (255)"`
	Night          int             `json:"night" gorm:"type : varchar (255)"`
	DateTrip       string          `json:"date_trip" gorm:"type : varchar (255)"`
	Price          int             `json:"price" gorm:"type : int"`
	Quota          int             `json:"quota" gorm:"type : int"`
	Description    string          `json:"description" gorm:"type : varchar (255)"`
	Image          string          `json:"image" gorm:"type : varchar (255)"`
}

type TripResponse struct {
	ID             int             `json:"id"`
	Title          string          `json:"title"`
	Country_id     int             `json:"country_id"`
	Country        CountryResponse `json:"country"`
	Accomodation   string          `json:"accomodation"`
	Transportation string          `json:"transportation"`
	Eat            string          `json:"eat"`
	Day            int             `json:"day"`
	Night          int             `json:"night"`
	DateTrip       string          `json:"date_trip"`
	Price          int             `json:"price"`
	Quota          int             `json:"quota"`
	Description    string          `json:"description"`
	Image          string          `json:"image"`
}

func (TripResponse) TableName() string {
	return "trips"
}
