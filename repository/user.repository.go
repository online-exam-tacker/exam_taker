package repository

import (
	// "errors"
	// "hamideh/data/request"
	"hamideh/helper"
	"hamideh/model"

	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	Db *gorm.DB
}

type UserRepository interface {
	Save(tags model.User)
	// Update(tags model.Tags)
	// Delete(tagsId int)
	// FindById(tagsId int) (tags model.Tags, err error)
	// FindAll() []model.Tags
}

func NewUserREpositoryImpl(Db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{Db: Db}
}

func (t *UserRepositoryImpl) Save(user model.User) {
	result := t.Db.Create(&user)
	helper.ErrorPanic(result.Error)
}

// Delete implements TagsRepository
// func (t *TagsRepositoryImpl) Delete(tagsId int) {
// 	var tags model.Tags
// 	result := t.Db.Where("id = ?", tagsId).Delete(&tags)
// 	helper.ErrorPanic(result.Error)
// }
