package service

import (
	"hamideh/data/request"
	"hamideh/data/response"
	"hamideh/helper"
	"hamideh/model"
	"hamideh/repository"

	"github.com/go-playground/validator/v10"
)

type ExamsService interface {
	Create(exams request.CreateExamRequest)
	Update(exams request.UpdateExamRequest)
	Delete(examsId int)
	FindById(examsId int) response.ExamResponse
	FindAll() []response.ExamResponse
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
		Name:     exams.Name,
		Question: question, // Assign the question to the exam
	}
	t.ExamsRepository.Save(examModel)
}

func (t *ExamsServiceImpl) Delete(examsId int) {
	t.ExamsRepository.Delete(examsId)
}

func (t *ExamsServiceImpl) FindAll() []response.ExamResponse {
	// Retrieve exam data from repository
	examData := t.ExamsRepository.FindAll()

	// Convert repository data to response format
	var exams []response.ExamResponse
	for _, exam := range examData {
		var questions []response.Question
		for _, q := range exam.Question {
			var responses []response.ModelResponse
			for _, r := range q.Responses {
				responses = append(responses, response.ModelResponse{
					Response: r.Response,
					IsTrue:   r.IsTrue,
				})
			}
			questions = append(questions, response.Question{
				Title:     q.Title,
				Responses: [4]response.ModelResponse(responses),
			})
		}
		exams = append(exams, response.ExamResponse{
			Id:       int(exam.ExamID),
			Name:     exam.Name,
			Question: questions, // Assuming the field in ExamResponse is Questions
		})
	}
	return exams
}

func (t *ExamsServiceImpl) FindById(examsId int) response.ExamResponse {
	examData, err := t.ExamsRepository.FindById(examsId)
	helper.ErrorPanic(err)

	var questions []response.Question
	for _, q := range examData.Question {
		var responses []response.ModelResponse
		for _, r := range q.Responses {
			responses = append(responses, response.ModelResponse{
				Response: r.Response,
				IsTrue:   r.IsTrue,
			})
		}
		questions = append(questions, response.Question{
			Title:     q.Title,
			Responses: [4]response.ModelResponse(responses),
		})
	}

	examResponse := response.ExamResponse{
		Id:       int(examData.ExamID),
		Name:     examData.Name,
		Question: questions,
	}
	return examResponse
}

func (t *ExamsServiceImpl) Update(examsId request.UpdateExamRequest) {
	examData, err := t.ExamsRepository.FindById(examsId.Id)
	helper.ErrorPanic(err)
	examData.Name = examsId.Name

	t.ExamsRepository.Update(examData)
}
