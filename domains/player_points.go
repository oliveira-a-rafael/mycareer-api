package domains

import (
	"errors"

	"github.com/jinzhu/gorm"
	"github.com/oliveira-a-rafael/mycareer-api/database"
	log "github.com/sirupsen/logrus"
)

type PlayerPoints struct {
	gorm.Model
	Player   Player `gorm:"foreignkey:PlayerID" json:"-"`
	PlayerID uint   `json:"player_id"`
	Points   int    `gorm:"index:player_points_idx" json:"points"`
	Age      int    `gorm:"index:player_age_idx" json:"age"`
}

func (playerPoints *PlayerPoints) Create() (*PlayerPoints, error) {

	if err := playerPoints.Validate(); err != nil {
		return playerPoints, err
	}

	db := database.GetInstance()
	err := db.Create(playerPoints).Error
	if err != nil {
		defer db.Close()
		return nil, err
	}

	err = db.Model(&Player{}).Where("id = ?", playerPoints.PlayerID).Update(Player{PointsCurrent: playerPoints.Points, AgeCurrent: playerPoints.Age}).Error
	if err != nil {
		defer db.Close()
		return nil, err
	}

	defer db.Close()
	return playerPoints, err
}

func (playerPoints *PlayerPoints) Validate() (err error) {
	player := playerPoints.Player
	player.ID = playerPoints.PlayerID

	db := database.GetInstance()
	err = db.First(&player, player.ID).Error
	db.Close()
	if err != nil {
		return err
	}

	if err := player.belongsToCareerAndAcoount(); err != nil {
		return err
	}

	var minPoints = 1
	var maxPoints = 100
	if playerPoints.Points < minPoints && playerPoints.Points > maxPoints {
		err = errors.New("points needs 1 between 100")
		log.Error(err)
		return err
	}

	return err
}
