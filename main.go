package main

import (
	"hamideh/config"
	"hamideh/controller"
	_ "hamideh/docs"
	"hamideh/helper"
	"hamideh/model"
	"hamideh/repository"
	"hamideh/router"
	"hamideh/service"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
)

// @title 	Tag Service API
// @version	1.0
// @description A Tag service API in Go using Gin framework

// @host 	localhost:8888
// @BasePath /api
func main() {

	log.Info().Msg("Started Server!")
	// Database
	db := config.DatabaseConnection()
	validate := validator.New()

	//Tables
	db.Table("exams").AutoMigrate(&model.Exam{})
	db.Table("responses").AutoMigrate(&model.Response{})
	db.Table("question").AutoMigrate(&model.Question{})

	// Repository
	// tagsRepository := repository.NewTagsREpositoryImpl(db)
	examRepositpry := repository.NewExamsREpositoryImpl(db)

	// Service
	// tagsService := service.NewTagsServiceImpl(tagsRepository, validate)
	examsService := service.NewExamServiceImpl(examRepositpry, validate)

	// Controller
	// tagsController := controller.NewTagsController(tagsService)
	examsController := controller.NewExamsController(examsService)

	// Router
	// routes := router.ExamRouter(tagsController)
	routes := router.ExamRouter(examsController)
	server := &http.Server{
		Addr:    ":8888",
		Handler: routes,
	}

	err := server.ListenAndServe()
	helper.ErrorPanic(err)
}
