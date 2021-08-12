package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/oliveira-a-rafael/mycareer-api/domains"
	"github.com/oliveira-a-rafael/mycareer-api/services"
	"github.com/oliveira-a-rafael/mycareer-api/utils"
)

type Accounter interface {
	CreateAccountNew(w http.ResponseWriter, r *http.Request)
}
type AccountController struct {
	// BaseController
	AccountService services.AccountContract
}

//func (c *AccountController) FindAccountByEmail(ctx *gin.Context) {
func (c *AccountController) CreateAccountNew(w http.ResponseWriter, r *http.Request) {

	account := &domains.Account{}
	err := json.NewDecoder(r.Body).Decode(account)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		utils.RespondError(w, err)
		return
	}

	fmt.Println("como e service: ", c.AccountService)

	// if err := c.AccountService.CreateAccount(account); err != nil {
	// 	w.WriteHeader(http.StatusUnprocessableEntity)
	// 	utils.RespondError(w, err)
	// 	return
	// }
}

// var CreateAccountNew = func(w http.ResponseWriter, r *http.Request) {
// 	account := &domains.Account{}
// 	err := json.NewDecoder(r.Body).Decode(account)
// 	if err != nil {
// 		w.WriteHeader(http.StatusUnprocessableEntity)
// 		utils.RespondError(w, err)
// 		return
// 	}

// 	w.WriteHeader(http.StatusCreated)
// 	utils.RespondNew(w, account)
// }

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
