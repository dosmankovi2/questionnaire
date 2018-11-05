package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/dosmankovi2/questionnaire/models"
	u "github.com/dosmankovi2/questionnaire/utils"
	"github.com/gorilla/mux"
)

//CreateQuestionnaire ...
var CreateQuestionnaire = func(w http.ResponseWriter, r *http.Request) {

	questionnaire := &models.Questionnaire{}

	err := json.NewDecoder(r.Body).Decode(questionnaire)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	resp := questionnaire.Create()
	u.Respond(w, resp)
}

//UpdateQuestionnaire ...
var UpdateQuestionnaire = func(w http.ResponseWriter, r *http.Request) {

	questionnaire := &models.Questionnaire{}

	err := json.NewDecoder(r.Body).Decode(questionnaire)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	resp := questionnaire.Update()
	u.Respond(w, resp)
}

//GetQuestionnaire ...
var GetQuestionnaire = func(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		u.Respond(w, u.Message(false, "There was an error in your request"))
		return
	}

	data := models.GetQuestionnaire(uint(id))
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

//DeleteQuestionnaire ...
var DeleteQuestionnaire = func(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		u.Respond(w, u.Message(false, "There was an error in your request"))
		return
	}

	data := models.DeleteQuestionnaire(uint(id))
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}
