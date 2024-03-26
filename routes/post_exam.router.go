package routes

import (
	"github.com/gorilla/mux"
	"github.com/hamideh/go_take_exam/controllers"
)

var RegisterExamTaker = func(router *mux.Router) {
	router.HandleFunc("/exam", controllers.PostExam).Methods("POST")
}
