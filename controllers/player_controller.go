package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/oliveira-a-rafael/mycareer-api/domains"
	"github.com/oliveira-a-rafael/mycareer-api/utils"
)

var CreatePlayer = func(w http.ResponseWriter, r *http.Request) {

	_, err := getLoggedUser(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		utils.RespondError(w, err)
		return
	}

	player := &domains.Player{}
	err = json.NewDecoder(r.Body).Decode(player)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		utils.RespondError(w, err)
		return
	}

	player, err = player.Create()
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		utils.RespondError(w, err)
		return
	}

	utils.RespondNew(w, player)
}

var GetPlayer = func(w http.ResponseWriter, r *http.Request) {

	//@TODO veriicar account
	id, err := GetID(w, r)

	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		utils.RespondError(w, err)
		return
	}

	player, err := domains.GetByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		utils.RespondError(w, err)
		return
	}

	utils.RespondNew(w, player)
}

var UpdatePlayer = func(w http.ResponseWriter, r *http.Request) {

	_, err := getLoggedUser(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		utils.RespondError(w, err)
		return
	}

	id, err := GetID(w, r)

	player := &domains.Player{}
	player.ID = id
	err = json.NewDecoder(r.Body).Decode(player)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		utils.RespondError(w, err)
		return
	}

	player, err = player.Update()
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		utils.RespondError(w, err)
		return
	}

	utils.RespondNew(w, player)
}

var DeletePlayer = func(w http.ResponseWriter, r *http.Request) {
	_, err := getLoggedUser(r)
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

	player := &domains.Player{}
	player.ID = id
	err = player.Delete()
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		utils.RespondError(w, err)
		return
	}

	utils.RespondNew(w, utils.MessageNew("deleted"))
}

var CreatePlayerPoints = func(w http.ResponseWriter, r *http.Request) {

	_, err := getLoggedUser(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		utils.RespondError(w, err)
		return
	}

	playerPoints := &domains.PlayerPoints{}
	err = json.NewDecoder(r.Body).Decode(playerPoints)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		utils.RespondError(w, err)
		return
	}

	playerPoints, err = playerPoints.Create()
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		utils.RespondError(w, err)
		return
	}

}

var ListPlayer = func(w http.ResponseWriter, r *http.Request) {

	player := &domains.Player{}
	lista, err := player.List()
	if err != nil {
		utils.RespondError(w, err)
		return
	}

	utils.RespondNew(w, lista)
}

var GetPosition = func(w http.ResponseWriter, r *http.Request) {

	positions := &domains.Positions

	positions_struct := []domains.Position{}

	for k, v := range *positions {
		positions_struct = append(positions_struct, domains.Position{
			Value:       k,
			Description: v,
		})
	}

	utils.RespondNew(w, positions_struct)
}
