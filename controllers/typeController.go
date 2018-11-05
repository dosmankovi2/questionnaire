package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/dosmankovi2/questionnaire/models"
	u "github.com/dosmankovi2/questionnaire/utils"
	"github.com/gorilla/mux"
)

//CreateQuestionype ...
var CreateQuestionType = func(w http.ResponseWriter, r *http.Request) {

	questionType := &models.QuestionType{}

	err := json.NewDecoder(r.Body).Decode(questionType)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	resp := questionType.Create()
	u.Respond(w, resp)
}

//GetQuestionType ...
var GetQuestionType = func(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		u.Respond(w, u.Message(false, "There was an error in your request"))
		return
	}

	data := models.GetQuestionType(uint(id))
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

//UpdateQuestionType ...
var UpdateQuestionType = func(w http.ResponseWriter, r *http.Request) {

	questionType := &models.QuestionType{}

	err := json.NewDecoder(r.Body).Decode(questionType)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	resp := questionType.Update()
	u.Respond(w, resp)
}

//DeleteQuestionType ...
var DeleteQuestionType = func(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		u.Respond(w, u.Message(false, "There was an error in your request"))
		return
	}

	data := models.DeleteQuestionType(uint(id))
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}
