package request

type CreateTagsRequest struct {
	Name string `validate:"required,min=1,max=200" json:"name"`
}

type UpdateTagsRequest struct {
	Id   int    `validate:"required"`
	Name string `validate:"required,max=200,min=1" json:"name"`
}

type Type struct {
	Four_option_exam string
	One_option_exam  string
}

type Question struct {
	Question_title string
	Responses      [4]Response
}

type Response struct {
	Response string
	Is_true  bool
}
type CreateExamRequest struct {
	Name      string     `validate:"required"`
	Type      Type       `validate:"required"`
	Questions []Question `validate:"required"`
}
