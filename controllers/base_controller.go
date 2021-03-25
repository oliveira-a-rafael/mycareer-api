package controllers

import (
	"errors"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
	"github.com/oliveira-a-rafael/my-career-api/utils"
)

type CareerTest struct {
	Title string
	Club  string
}

var ListaCareersToTest = func(w http.ResponseWriter, r *http.Request) {
	careers := []CareerTest{}

	pb := CareerTest{
		Title: "Career One",
		Club:  "Liverpoll",
	}

	careers = append(careers, pb)

	utils.RespondNew(w, careers)
}

var Cors = func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method == http.MethodOptions {
		return
	}
}

var HealthCheck = func(w http.ResponseWriter, r *http.Request) {
	utils.RespondNew(w, "up")
}

var getLoggedUser = func(r *http.Request) (uint, error) {
	userLogged := r.Context().Value("user")

	if userLogged == nil {
		err := errors.New("error on access user logged info")
		log.Println(err.Error())
		return 0, err
	}

	userId := userLogged.(uint)
	userId, ok := userLogged.(uint)
	if !ok {
		err := errors.New("error on convert user logged info")
		log.Println(err.Error())
		return 0, err
	}

	return userId, nil
}

var GetID = func(w http.ResponseWriter, r *http.Request) (uint, error) {
	params := mux.Vars(r)
	if params["id"] == "" {
		err := errors.New("ID entity not found")
		utils.Respond(w, utils.Message(false, err.Error()))
		return 0, err
	}

	id, err := utils.ConvertToUint(params["id"])
	if err != nil {
		log.Error(err)
		utils.Respond(w, utils.Message(false, err.Error()))
		return 0, err
	}

	return id, nil
}
