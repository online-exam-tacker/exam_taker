package service

import (
	"hamideh/data/request"
	"time"

	// "hamideh/data/response"
	"hamideh/helper"
	"hamideh/model"
	"hamideh/repository"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-playground/validator/v10"
)

type UsersService interface {
	Create(users request.CreateUserRequest)
	Payload(usersRequest request.CreateUserRequest) string
	// Update(tags request.UpdateTagsRequest)
	// Delete(tagsId int)
	// FindById(tagsId int) response.TagsResponse
	// FindAll() []response.TagsResponse
}

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	Validate       *validator.Validate
}

func NewUserServiceImpl(userRepository repository.UserRepository, validate *validator.Validate) UsersService {
	return &UserServiceImpl{
		UserRepository: userRepository,
		Validate:       validate,
	}
}

var jwtKey = []byte("your_secret_key")

func (t *UserServiceImpl) Create(users request.CreateUserRequest) {
	err := t.Validate.Struct(users)
	helper.ErrorPanic(err)
	userModel := model.User{
		Username: users.Username,
		Password: users.Password,
		Role:     "user",
	}
	t.UserRepository.Save(userModel)
}

func (t *UserServiceImpl) Payload(usersRequest request.CreateUserRequest) string {
	err := t.Validate.Struct(usersRequest)
	helper.ErrorPanic(err)

	userName, err := t.UserRepository.FindById(usersRequest.Id)
	helper.ErrorPanic(err)

	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &request.CreateUserRequest{
		Username: userName.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		helper.ErrorPanic(err)
	}
	return tokenString
}
