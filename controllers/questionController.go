package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/dosmankovi2/questionnaire/models"
	u "github.com/dosmankovi2/questionnaire/utils"
	"github.com/gorilla/mux"
)

//CreateQuestion ...
var CreateQuestion = func(w http.ResponseWriter, r *http.Request) {

	question := &models.Question{}

	err := json.NewDecoder(r.Body).Decode(question)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	resp := question.Create()
	u.Respond(w, resp)
}

//GetQuestion ...
var GetQuestion = func(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		u.Respond(w, u.Message(false, "There was an error in your request"))
		return
	}

	data := models.GetQuestion(uint(id))
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

//UpdateQuestion ...
var UpdateQuestion = func(w http.ResponseWriter, r *http.Request) {

	question := &models.Question{}

	err := json.NewDecoder(r.Body).Decode(question)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	resp := question.Update()
	u.Respond(w, resp)
}

//DeleteQuestion ...
var DeleteQuestion = func(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		u.Respond(w, u.Message(false, "There was an error in your request"))
		return
	}

	data := models.DeleteQuestion(uint(id))
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}
