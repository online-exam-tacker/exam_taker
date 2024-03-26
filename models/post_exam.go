package models

import (
	"log"

	"github.com/hamideh/go_take_exam/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:"" json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

type Exam struct {
	gorm.Model
	id         string     `gorm:"" json:"id"`
	questions  Questions  `json:"questions"`
	collection Collection `json:"collection"`
	answer     Options    `json:"answer"`
}

type Questions struct {
	gorm.Model
	id       string  `json:"id"`
	question string  `json:"question"`
	answer   Options `json:"answer"`
}

type Options struct {
	gorm.Model
	op1 string `json:"op1"`
	op2 string `json:"op2"`
	op3 string `json:"op3"`
	op4 string `json:"op4"`
}

type Collection struct {
	gorm.Model
	id   string `gorm:"" json:"id"`
	name string ` json:"name"`
}

func (e *Exam) CreateExam() *Exam {
	exam := Exam{questions: e.questions, collection: e.collection, answer: e.answer}

	result := db.Create(&exam)
	if result.Error != nil {
		log.Fatalln(result.Error)
	}
	return e
}

// func (e *Exam) CreateExam() *Exam {
// 	db.NewRecord(e)
// 	db.Create(&e)
// 	return e
// }

func GetExams() []Exam {
	var Exams []Exam
	db.Find(&Exams)
	return Exams
}

func GetExamById(Id int64) (*Exam, *gorm.DB) {
	var getExam Exam
	db := db.Where("ID=?", Id).Find(&getExam)
	return &getExam, db
}

func DeleteBook(ID int64) Exam {
	var Exam Exam
	db.Where("ID=?", ID).Delete(Exam)
	return Exam
}
