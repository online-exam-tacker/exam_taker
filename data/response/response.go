package response

type TagsResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Response struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data,omitempty"`
}

type ModelResponse struct {
	Response string
	IsTrue   bool
}

type Question struct {
	Title     string
	Responses [4]ModelResponse `validate:"required"`
}

type ExamResponse struct {
	Name     string
	Question []Question
}
