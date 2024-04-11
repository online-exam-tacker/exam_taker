package request

type CreateTagsRequest struct {
	Name string `validate:"required,min=1,max=200" json:"name"`
}

type UpdateTagsRequest struct {
	Id   int    `validate:"required"`
	Name string `validate:"required,max=200,min=1" json:"name"`
}

type Question struct {
	Title     string
	Responses [4]Response `validate:"required"`
}

type Response struct {
	Response string
	IsTrue   bool
}

type CreateExamRequest struct {
	Name     string
	Question []Question
}
