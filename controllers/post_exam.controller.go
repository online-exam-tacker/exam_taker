package controllers

import (
	"net/http"

	"github.com/hamideh/go_take_exam/models"

	"encoding/json"
	"io"
)

func ParseBody(r *http.Request, x interface{}) {
	if body, err := io.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}
}

func PostExam(w http.ResponseWriter, r *http.Request) {

	exam_model := &models.Exam{}
	ParseBody(r, exam_model)
	e := exam_model.CreateExam()

	res, _ := json.Marshal(e)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
