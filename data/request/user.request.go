package request

import "github.com/dgrijalva/jwt-go"

type CreateUserRequest struct {
	Id       int
	Username string `validate:"required,max=200,min=1" json:"username"`
	Password string `validate:"required,max=200,min=4" json:"password"`
	Role     string `validate:"required"`
	jwt.StandardClaims
}
