package service

import (
	"hamideh/data/request"
	// "hamideh/data/response"
	"hamideh/helper"
	"hamideh/model"
	"hamideh/repository"

	"github.com/go-playground/validator/v10"
)

type ExamsService interface {
	Create(exams request.CreateExamRequest)
	// Update(tags request.UpdateTagsRequest)
	// Delete(tagsId int)
	// FindById(tagsId int) response.TagsResponse
	// FindAll() []response.TagsResponse
}

type ExamsServiceImpl struct {
	ExamsRepository repository.ExamsRepository
	Validate        *validator.Validate
}

func NewExamServiceImpl(ExamsRepository repository.ExamsRepository, validate *validator.Validate) ExamsService {
	return &ExamsServiceImpl{
		ExamsRepository: ExamsRepository,
		Validate:        validate,
	}
}

// Create implements TagsService
func (t *ExamsServiceImpl) Create(exams request.CreateExamRequest) {
	err := t.Validate.Struct(exams)
	helper.ErrorPanic(err)

	var responses []model.Response
	for _, q := range exams.Question {
		for _, r := range q.Responses {
			responses = append(responses, model.Response{
				Response: r.Response,
				IsTrue:   r.IsTrue,
			})
		}
	}

	var question []model.Question
	for _, q := range exams.Question {
		question = append(question, model.Question{
			Title:     q.Title,
			Responses: responses, // Assign the responses to each question
		})
	}

	examModel := model.Exam{
		Name: exams.Name,

		Question: question, // Assign the question to the exam
	}

	t.ExamsRepository.Save(examModel)
}

// Delete implements TagsService
// func (t *TagsServiceImpl) Delete(tagsId int) {
// 	t.TagsRepository.Delete(tagsId)
// }

// // FindAll implements TagsService
// func (t *TagsServiceImpl) FindAll() []response.TagsResponse {
// 	result := t.TagsRepository.FindAll()

// 	var tags []response.TagsResponse
// 	for _, value := range result {
// 		tag := response.TagsResponse{
// 			Id:   value.Id,
// 			Name: value.Name,
// 		}
// 		tags = append(tags, tag)
// 	}

// 	return tags
// }

// // FindById implements TagsService
// func (t *TagsServiceImpl) FindById(tagsId int) response.TagsResponse {
// 	tagData, err := t.TagsRepository.FindById(tagsId)
// 	helper.ErrorPanic(err)

// 	tagResponse := response.TagsResponse{
// 		Id:   tagData.Id,
// 		Name: tagData.Name,
// 	}
// 	return tagResponse
// }

// // Update implements TagsService
// func (t *TagsServiceImpl) Update(tags request.UpdateTagsRequest) {
// 	tagData, err := t.TagsRepository.FindById(tags.Id)
// 	helper.ErrorPanic(err)
// 	tagData.Name = tags.Name
// 	t.TagsRepository.Update(tagData)
// }
