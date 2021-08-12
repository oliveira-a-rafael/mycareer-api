package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/oliveira-a-rafael/mycareer-api/domains"
	"github.com/oliveira-a-rafael/mycareer-api/services"
	"github.com/oliveira-a-rafael/mycareer-api/utils"
)

type AccountHandler interface {
	GetAccount(w http.ResponseWriter, r *http.Request)
	CreateAccount(w http.ResponseWriter, r *http.Request)
}

type AccountController struct {
	AccountService services.AccountService
}

func (ac *AccountController) CreateAccount(w http.ResponseWriter, r *http.Request) {

	account := &domains.Account{}
	err := json.NewDecoder(r.Body).Decode(account)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	account, err = ac.AccountService.CreateAccount(account)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		utils.RespondError(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	utils.RespondNew(w, account)

}
