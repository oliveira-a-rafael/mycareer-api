package domains

import (
	"encoding/json"
	"errors"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/oliveira-a-rafael/mycareer-api/database"
	log "github.com/sirupsen/logrus"
)

type Career struct {
	gorm.Model
	Account     Account  `gorm:"foreignkey:AccountID; not null" json:"-"`
	AccountID   uint     `json:"-"`
	Title       string   `json:"title"`
	TeamName    string   `json:"team_name"`
	Players     []Player `gorm:"foreignkey:CareerID" json:"players"`
	TotalPlayer int      `gorm:"-" json:"total_players"`
}

func (career *Career) Create() (*Career, error) {
	db := database.GetInstance()

	if err := career.Validate(); err != nil {
		return career, err
	}

	err := db.Create(career).Error
	defer db.Close()

	if err != nil {
		return career, err
	}

	return career, nil
}

func (career *Career) Update() (*Career, error) {

	if err := career.Validate(); err != nil {
		return career, err
	}

	if err := career.belongsToAccount(); err != nil {
		return career, err
	}

	db := database.GetInstance()
	err := db.Save(career).Error
	defer db.Close()

	if err != nil {
		return career, err
	}

	return career, err
}

func (career *Career) Delete() (err error) {

	db := database.GetInstance()
	var careerAccount Career
	db.Find(&careerAccount, "id = ? and account_id = ?", career.ID, career.AccountID)
	defer db.Close()

	if careerAccount.ID < 1 {
		err = errors.New("career do not belongs to this account")
		return err
	}

	db = database.GetInstance()
	err = db.Delete(&career).Error
	defer db.Close()
	if err != nil {
		return err
	}

	return err
}

func (career *Career) Validate() (err error) {

	if strings.TrimSpace(career.Title) == "" {
		err = errors.New("title is required")
		log.Error(err)
		return err
	}

	if strings.TrimSpace(career.TeamName) == "" {
		err = errors.New("team name is required")
		log.Error(err)
		return err
	}

	return err
}

func (career *Career) Parse() map[string]interface{} {
	var data map[string]interface{}
	in, err := json.Marshal(&career)
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(in, &data); err != nil {
		log.Error(err.Error())
	}
	return data
}

func (career *Career) GetCareer() (*Career, error) {

	db := database.GetInstance()

	err := db.Find(&career, "id = ? and account_id = ?", career.ID, career.AccountID).Error
	if err != nil {
		defer db.Close()
		return career, err
	}

	err = db.Model(&Player{}).Where("career_id = ?", career.ID).Count(&career.TotalPlayer).Error
	if err != nil {
		defer db.Close()
		return career, err
	}

	defer db.Close()
	return career, err
}

func (career *Career) belongsToAccount() (err error) {

	db := database.GetInstance()

	found := Career{}
	err = db.Find(&found, "id = ? and account_id = ?", career.ID, career.AccountID).Error
	if err != nil {
		defer db.Close()
		return errors.New("Error on get carrer for this account")
	}

	if found.ID < 1 || found.AccountID < 1 {
		defer db.Close()
		return errors.New("Error on get carrer for this account")
	}

	defer db.Close()
	return nil
}

func (career *Career) ListPlayers() ([]Player, error) {

	err := career.belongsToAccount()
	if err != nil {
		return nil, err
	}

	var players []Player

	db := database.GetInstance()
	err = db.First(&career, career.ID).Related(&players).Error
	defer db.Close()
	if err != nil {
		return nil, err
	}

	return players, err

}
