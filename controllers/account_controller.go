package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/oliveira-a-rafael/mycareer-api/domains"
	"github.com/oliveira-a-rafael/mycareer-api/services"
	"github.com/oliveira-a-rafael/mycareer-api/utils"
)

type AccountController struct {
	AccountService services.AccountContract
}

var CreateAccountNew = func(w http.ResponseWriter, r *http.Request) {
	account := &domains.Account{}
	err := json.NewDecoder(r.Body).Decode(account)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		utils.RespondError(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	utils.RespondNew(w, account)
}

// var CreateAccount = func(w http.ResponseWriter, r *http.Request) {

// 	// account := &domains.Account{}
// 	// err := json.NewDecoder(r.Body).Decode(account)
// 	// if err != nil {
// 	// 	w.WriteHeader(http.StatusUnprocessableEntity)
// 	// 	utils.RespondError(w, err)
// 	// 	return
// 	// }

// 	// account, err = account.Create()
// 	// if err != nil {
// 	// 	w.WriteHeader(http.StatusUnprocessableEntity)
// 	// 	utils.RespondError(w, err)
// 	// 	return
// 	// }

// 	// w.WriteHeader(http.StatusCreated)
// 	utils.RespondNew(w, account)
// }
