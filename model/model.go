package model

import (
	"gorm.io/gorm"
)

type Tags struct {
	Id   int    `gorm:"type:int;primary_key"`
	Name string `gorm:"type:varchar(255)"`
}

type User struct {
	gorm.Model
	Id       int    `gorm:"type:int;primary_key"`
	Exams    []Exam `gorm:"foreignKey:Exam_ID"`
	Password string
}

type Type struct {
	Four_option_exam string
	One_option_exam  string
}

type Exam struct {
	Exam_ID   uint `gorm:"primary_key"`
	Type      Type
	Questions []Question `gorm:"foreignKey:Question_ID"`
	Responses []Response `gorm:"foreignKey:Response_ID"`
	Name      string
}

type Response struct {
	Response_ID uint
	Response    string
	Is_true     bool
}

type Question struct {
	Question_ID    uint `gorm:"primary_key"`
	Answer_ID      int  `gorm:"foreignKey:Answer_ID"`
	Question_title string
	Responses      [4]Response
}

// type Answer struct {
// 	Answer_ID uint `gorm:"primary_key"`
// 	Answer    string
// }
