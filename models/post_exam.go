// package models

// import(
// 	"gorm.io/driver/postgres"
//     "gorm.io/gorm"
// )

// type User struct {
//     gorm.Model
//     Name string
//     Age  int
// }

//	func createUser(db *gorm.DB, name string, age int) error {
//	    user := User{Name: name, Age: age}
//	    result := db.Create(&user)
//	    if result.Error != nil {
//	        return result.Error
//	    }
//	    return nil
//	}
package models

import (
	"github.com/hamideh/go_take_exam/config"
	"gorm.io/gorm"
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

// func (b *Book) CreateBook() *Book {
// 	db.NewRecord(b)
// 	db.Create(&b)
// 	return b
// }

// func GetBooks() []Book {
// 	var Books []Book
// 	db.Find(&Books)
// 	return Books
// }

// func GetBookById(Id int64) (*Book, *gorm.DB) {
// 	var getBook Book
// 	db := db.Where("ID=?", Id).Find(&getBook)
// 	return &getBook, db
// }

// func DeleteBook(ID int64) Book {
// 	var book Book
// 	db.Where("ID=?", ID).Delete(book)
// 	return book
// }

type Exam struct {
	id         string
	questions  Questions
	collection Collection
	answer     Options
}

type Questions struct {
	question string
	answer   Options
}

type Options struct {
	op1 string
	op2 string
	op3 string
	op4 string
}

type Collection struct {
	id   string
	name string
}

func createExam(db *gorm.DB, questions Questions, collection Collection, answer Options) error {
	exam := Exam{questions: questions, collection: collection, answer: answer}

	result := db.Create(&exam)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
