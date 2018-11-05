package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/dosmankovi2/questionnaire/models"
	u "github.com/dosmankovi2/questionnaire/utils"
	"github.com/gorilla/mux"
)

//CreateChoice ...
var CreateChoice = func(w http.ResponseWriter, r *http.Request) {

	choice := &models.QuestionChoice{}

	err := json.NewDecoder(r.Body).Decode(choice)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	resp := choice.Create()
	u.Respond(w, resp)
}

//GetChoice ...
var GetChoice = func(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		u.Respond(w, u.Message(false, "There was an error in your request"))
		return
	}

	data := models.GetQuestionChoice(uint(id))
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

//UpdateChoice ...
var UpdateChoice = func(w http.ResponseWriter, r *http.Request) {

	choice := &models.QuestionChoice{}

	err := json.NewDecoder(r.Body).Decode(choice)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	resp := choice.Update()
	u.Respond(w, resp)
}

//DeleteChoice ...
var DeleteChoice = func(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		u.Respond(w, u.Message(false, "There was an error in your request"))
		return
	}

	data := models.DeleteQuestionChoice(uint(id))
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}
