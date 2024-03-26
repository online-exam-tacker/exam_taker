// package config

// import (
// 	"database/sql"
// 	"log"

// 	"gorm.io/driver/postgres"
// 	"gorm.io/gorm"
// )

// var (
// 	db *gorm.DB
// )

// func Connect() {
// 	// d, err := gorm.Open("postgres", "hamideh:alexa@12@/simplerest?charset=utf8&parseTime=True&loc=Local")
// 	// if err != nil {
// 	// 	panic(err)
// 	// }
// 	// db = d
// 	sqlDB, err := sql.Open("pgx", "examtaker_db")
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	d, err := gorm.Open(postgres.New(postgres.Config{
// 		Conn: sqlDB,
// 	}), &gorm.Config{})

// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	db = d
// }

//	func GetDB() *gorm.DB {
//		return db
//	}
package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	db *gorm.DB
)

func Connect() {
	// d, err := gorm.Open("postgres", "postgres://hamiltoon:1234@0.0.0.0:5432/examtaker_db")
	d, err := gorm.Open("postgresql", "postgresql://hamiltoon:1234@127.0.0.0:5430/examtaker_db")
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
