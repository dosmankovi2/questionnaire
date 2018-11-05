package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/dosmankovi2/questionnaire/models"
	util "github.com/dosmankovi2/questionnaire/utils"
)

//CreateAccount action method
var CreateAccount = func(w http.ResponseWriter, r *http.Request) {
	account := &models.Account{}
	err := json.NewDecoder(r.Body).Decode(account)

	if err != nil {
		util.Respond(w, util.Message(false, "Invalid request. Please try again."))
		return
	}

	resp := account.Create()
	util.Respond(w, resp)
}

//Authenticate action method
var Authenticate = func(w http.ResponseWriter, r *http.Request) {
	account := &models.Account{}
	err := json.NewDecoder(r.Body).Decode(account)

	if err != nil {
		util.Respond(w, util.Message(false, "Invalid request. Please try again."))
		return
	}

	resp := models.Login(account.Email, account.Password)
	util.Respond(w, resp)
}
