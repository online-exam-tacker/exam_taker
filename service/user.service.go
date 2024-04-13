package service

import (
	"hamideh/data/request"
	// "hamideh/data/response"
	"hamideh/helper"
	"hamideh/model"
	"hamideh/repository"

	"github.com/go-playground/validator/v10"
)

type UsersService interface {
	Create(users request.CreateUserRequest)
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

// Create implements TagsService
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
