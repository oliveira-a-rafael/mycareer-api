package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/oliveira-a-rafael/my-career-api/domains"
	"github.com/oliveira-a-rafael/my-career-api/utils"
)

var CreateAccount = func(w http.ResponseWriter, r *http.Request) {

	account := &domains.Account{}
	err := json.NewDecoder(r.Body).Decode(account)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		utils.RespondError(w, err)
		return
	}

	account, err = account.Create()
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		utils.RespondError(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	utils.RespondNew(w, account)
}

var Authenticate = func(w http.ResponseWriter, r *http.Request) {

	var account = &domains.Account{}
	err := json.NewDecoder(r.Body).Decode(account)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		utils.RespondError(w, err)
		return
	}

	account, err = domains.Login(account.Email, account.Password)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		utils.RespondError(w, err)
		return
	}

	utils.RespondNew(w, account)
}

var ListCareers = func(w http.ResponseWriter, r *http.Request) {

	userLogged, err := getLoggedUser(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		utils.RespondError(w, err)
		return
	}

	account := &domains.Account{}
	account.ID = userLogged
	err, careers := account.ListCareers()
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		utils.RespondError(w, err)
		return
	}

	utils.RespondNew(w, careers)
}
