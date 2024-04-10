package repository

import (
	"errors"
	"hamideh/data/request"
	"hamideh/helper"
	"hamideh/model"

	"gorm.io/gorm"
)

type ExamsRepositoryImpl struct {
	Db *gorm.DB
}

type ExamsRepository interface {
	Save(exams model.Exam)
	Update(exams model.Exam)
	Delete(examsId int)
	FindById(examsId int) (exams model.Exam, err error)
	FindAll() []model.Exam
}

func NewExamsREpositoryImpl(Db *gorm.DB) ExamsRepository {
	return &ExamsRepositoryImpl{Db: Db}
}

// // Save implements TagsRepository
func (t *ExamsRepositoryImpl) Save(exams model.Exam) {
	result := t.Db.Create(&exams)
	helper.ErrorPanic(result.Error)
}

// Delete implements TagsRepository
func (t *ExamsRepositoryImpl) Delete(examsId int) {
	var exams model.Exam
	result := t.Db.Where("id = ?", examsId).Delete(&exams)
	helper.ErrorPanic(result.Error)
}

// // FindAll implements TagsRepository
func (t *ExamsRepositoryImpl) FindAll() []model.Exam {
	var exams []model.Exam
	result := t.Db.Find(&exams)
	helper.ErrorPanic(result.Error)
	return exams
}

// // FindById implements TagsRepository
func (t *ExamsRepositoryImpl) FindById(examsId int) (exams model.Exam, err error) {
	var exam model.Exam
	result := t.Db.Find(&exam, examsId)
	if result != nil {
		return exam, nil
	} else {
		return exam, errors.New("tag is not found")
	}
}

// // Update implements TagsRepository
func (t *ExamsRepositoryImpl) Update(exams model.Exam) {
	var updateTag = request.UpdateTagsRequest{
		Id:   int(exams.ExamID),
		Name: exams.Name,
	}
	result := t.Db.Model(&exams).Updates(updateTag)
	helper.ErrorPanic(result.Error)
}
