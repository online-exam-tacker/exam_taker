package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (a *application) routes() http.Handler {

	router := chi.NewRouter()
	router.Post("/exam", a.PostExam)

}
