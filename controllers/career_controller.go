package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/oliveira-a-rafael/mycareer-api/domains"
	"github.com/oliveira-a-rafael/mycareer-api/utils"
)

var CreateCareer = func(w http.ResponseWriter, r *http.Request) {

	userLogged, err := getLoggedUser(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		utils.RespondError(w, err)
		return
	}

	career := &domains.Career{}
	career.AccountID = userLogged
	err = json.NewDecoder(r.Body).Decode(career)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		utils.RespondError(w, err)
		return
	}

	career, err = career.Create()
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		utils.RespondError(w, err)
		return
	}

	utils.RespondNew(w, career)
}

var UpdateCareer = func(w http.ResponseWriter, r *http.Request) {

	userLogged, err := getLoggedUser(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		utils.RespondError(w, err)
		return
	}

	id, err := GetID(w, r)

	career := &domains.Career{}
	career.ID = id
	career.AccountID = userLogged
	err = json.NewDecoder(r.Body).Decode(career)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		utils.RespondError(w, err)
		return
	}

	career, err = career.Update()
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		utils.RespondError(w, err)
		return
	}

	utils.RespondNew(w, career)
}

var DeleteCareer = func(w http.ResponseWriter, r *http.Request) {
	userLogged, err := getLoggedUser(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		utils.RespondError(w, err)
		return
	}

	id, err := GetID(w, r)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		utils.RespondError(w, err)
		return
	}

	career := &domains.Career{}
	career.ID = id
	career.AccountID = userLogged
	err = career.Delete()
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		utils.RespondError(w, err)
		return
	}

	utils.RespondNew(w, utils.MessageNew("deleted"))
}

var GetCareer = func(w http.ResponseWriter, r *http.Request) {
	userLogged, err := getLoggedUser(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		utils.RespondError(w, err)
		return
	}

	id, err := GetID(w, r)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		utils.RespondError(w, err)
		return
	}

	career := &domains.Career{}
	career.ID = id
	career.AccountID = userLogged
	career, err = career.GetCareer()

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		utils.RespondError(w, err)
		return
	}

	utils.RespondNew(w, career)
}

var CareerPlayers = func(w http.ResponseWriter, r *http.Request) {
	userLogged, err := getLoggedUser(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		utils.RespondError(w, err)
		return
	}

	id, err := GetID(w, r)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		utils.RespondError(w, err)
		return
	}

	career := &domains.Career{}
	career.ID = id
	career.AccountID = userLogged
	players, err := career.ListPlayers()

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		utils.RespondError(w, err)
		return
	}

	utils.RespondNew(w, players)
}
