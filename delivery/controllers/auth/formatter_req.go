package auth

type LoginRequestFormat struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}
type AdminRequestFormat struct {
	Email        string `json:"email" form:"email"`
	Password     string `json:"password" form:"password"`
	Name         string `json:"name" form:"name"`
	Phone_Number string `json:"phone_number" form:"phone_number"`
}
