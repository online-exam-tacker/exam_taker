package repository

import (
	"errors"
	// "hamideh/data/request"
	"hamideh/helper"
	"hamideh/model"

	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	Db *gorm.DB
}

type UserRepository interface {
	Save(users model.User)
	// GetUsername(username string)
	// Update(tags model.Tags)
	// Delete(tagsId int)
	FindById(usersId int) (users model.User, err error)
	// FindAll() []model.Tags
}

func NewUserREpositoryImpl(Db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{Db: Db}
}

func (t *UserRepositoryImpl) Save(user model.User) {
	result := t.Db.Create(&user)
	helper.ErrorPanic(result.Error)
}

func (t *UserRepositoryImpl) FindById(userId int) (users model.User, err error) {
	var user model.User
	result := t.Db.Find(&user, userId)
	if result != nil {
		return user, nil
	} else {
		return user, errors.New("user is not found")
	}
}

// Delete implements TagsRepository
// func (t *TagsRepositoryImpl) Delete(tagsId int) {
// 	var tags model.Tags
// 	result := t.Db.Where("id = ?", tagsId).Delete(&tags)
// 	helper.ErrorPanic(result.Error)
// }
