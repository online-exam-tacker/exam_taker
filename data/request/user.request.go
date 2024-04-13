package request

type CreateUserRequest struct {
	Username string `validate:"required,max=200,min=1" json:"username"`
	Password string `validate:"required,max=200,min=4" json:"password"`
	Role     string `validate:"required"`
}
