package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/dosmankovi2/questionnaire/models"
	u "github.com/dosmankovi2/questionnaire/utils"
	"github.com/gorilla/mux"
)

//CreateAnswer ...
var CreateAnswer = func(w http.ResponseWriter, r *http.Request) {

	answer := &models.Answer{}

	err := json.NewDecoder(r.Body).Decode(answer)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	resp := answer.Create()
	u.Respond(w, resp)
}

//GetAnswer ...
var GetAnswer = func(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		u.Respond(w, u.Message(false, "There was an error in your request"))
		return
	}

	data := models.GetAnswer(uint(id))
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

//UpdateAnswer ...
var UpdateAnswer = func(w http.ResponseWriter, r *http.Request) {

	answer := &models.Answer{}

	err := json.NewDecoder(r.Body).Decode(answer)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	resp := answer.Update()
	u.Respond(w, resp)
}

//DeleteAnswer ...
var DeleteAnswer = func(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		u.Respond(w, u.Message(false, "There was an error in your request"))
		return
	}

	data := models.DeleteAnswer(uint(id))
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}
