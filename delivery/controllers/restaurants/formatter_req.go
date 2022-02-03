package restaurants

import "time"

type LoginRequestFormat struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type RestaurantRequestFormat struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type CreateRestaurantDRequestFormat struct {
	Name           string    `json:"name"`
	Open           time.Time `json:"open"`
	Close          time.Time `json:"close"`
	Price          int       `json:"price"`
	Latitude       float64   `json:"latitude"`
	Longitude      float64   `json:"longitude"`
	City           string    `json:"city"`
	Address        string    `json:"address"`
	PhoneNumber    string    `json:"phone"`
	ProfilePicture string    `json:"profile_picture"`
	Seats          int       `json:"seats"`
	Description    string    `json:"description"`
}

type UpdateRestaurantDRequestFormat struct {
	Open           time.Time `json:"open"`
	Close          time.Time `json:"close"`
	Price          int       `json:"price"`
	PhoneNumber    string    `json:"phone"`
	ProfilePicture string    `json:"profile_picture"`
	Seats          int       `json:"json"`
	Description    string    `json:"description"`
}
