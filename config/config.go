package config

import (
	"database/sql"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Connect() {
	// d, err := gorm.Open("postgres", "hamideh:alexa@12@/simplerest?charset=utf8&parseTime=True&loc=Local")
	// if err != nil {
	// 	panic(err)
	// }
	// db = d
	sqlDB, err := sql.Open("pgx", "examtaker_db")
	if err != nil {
		log.Fatalln(err)
	}

	d, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db = d
}

func GetDB() *gorm.DB {
	return db
}
