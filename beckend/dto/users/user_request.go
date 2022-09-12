package usersdto

type CreateUserRequest struct {
	Name     string `json:"name" form:"name" validate:"required"`
	Email    string `json:"email" form:"email" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
}

type UpdateUserRequest struct {
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Address  string `json:"address" form:"address"`
	Postcode string `json:"postcode" form:"postcode"`
	Image    string `json:"image" form:"image"`
}
