package model

import (
	"github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
)

type Tags struct {
	Id   int    `gorm:"type:int;primary_key"`
	Name string `gorm:"type:varchar(255)"`
}

type User struct {
	gorm.Model
	Id       int    `gorm:"type:int;primary_key"`
	Exams    []Exam `gorm:"foreignKey:ExamID"`
	Password string
	Username string `json:"username"`
	Role     string
	jwt.StandardClaims
}

type Exam struct {
	ExamID   uint       `gorm:"primary_key"`
	Question []Question `gorm:"foreignKey:ExamID"`
	Name     string
}

type Response struct {
	ResponseID uint
	Response   string
	IsTrue     bool `gorm:"column:istrue"` // Corrected column name to match database
	QuestionID uint // Added foreign key field to establish the relationship
}

type Question struct {
	QuestionID uint `gorm:"primary_key"`
	Title      string
	Responses  []Response `gorm:"foreignKey:QuestionID"` // Specifying foreign key for Responses
	ExamID     uint       // Added foreign key field to establish the relationship
}

// type Exam struct {
// 	ExamID uint `gorm:"primary_key"`

// 	question []question `gorm:"foreignKey:QuestionID"`

// 	Name string
// }

// type Responses struct {
// 	ResponseID uint
// 	Response   string
// 	Is_true    bool
// }

// type question struct {
// 	QuestionID uint `gorm:"primary_key"`
// 	Title      string
// 	Responses  []Responses
// }
