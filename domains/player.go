package domains

import (
	"errors"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/oliveira-a-rafael/my-career-api/database"
	"github.com/oliveira-a-rafael/my-career-api/utils"
	log "github.com/sirupsen/logrus"
)

// Player model
type Player struct {
	gorm.Model
	Career           Career        `gorm:"foreignkey:CareerID" json:"-"`
	CareerID         uint          `json:"career_id"`
	Name             string        `gorm:"type:varchar(255);index:player_name_idx" json:"name"`
	Position         string        `gorm:"type:varchar(100)" json:"position"`
	AgeCurrent       int           `json:"age_current" json:"age_current"`
	AgeInitial       int           `json:"age_initial" json:"age_initial"`
	StrongFoot       string        `gorm:"type:varchar(100)" json:"strong_foot"`
	PointsInitial    int           `gorm:"index:points_init_idx" json:"points_initial"`
	PointsCurrent    int           `gorm:"index:points_current_idx" json:"points_current"`
	PotentialInitial int           `gorm:"index:pot_points_init_idx" json:"potential_initial"`
	PotentialFinal   int           `gorm:"index:pot_points_idx" json:"potential_final"`
	Skills           []PlayerSkill `gorm:"foreignkey:PlayerID" json:"skills"`
}

type SkillsResponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Value int    `json:"value"`
}

func (player *Player) belongsToCareerAndAcoount() (err error) {

	accountID, err := utils.GetCurrentUser()
	if err != nil {
		log.Error(err)
		return err
	}

	db := database.GetInstance()
	career := &Career{}
	err = db.Where("id = ? and account_id = ? ", player.CareerID, accountID).First(career).Error
	db.Close()
	if err != nil {
		log.Error(err)
		return err
	}

	if career.ID < 1 {
		err := errors.New("Error in account")
		log.Error(err)
		return err
	}

	return err
}

func (player *Player) Create() (*Player, error) {

	if err := player.Validate(); err != nil {
		return player, err
	}

	if player.PointsCurrent < 1 {
		player.PointsCurrent = player.PointsInitial
	}

	if player.AgeCurrent < 1 {
		player.AgeCurrent = player.AgeInitial
	}

	db := database.GetInstance()
	err := db.Create(player).Error
	defer db.Close()

	if err != nil {
		return nil, err
	}

	if err := saveHistory(player.ID); err != nil {
		return nil, err
	}

	return player, nil
}

func saveHistory(playerID uint) error {
	skills, err := GetSkills(playerID)
	if err != nil {
		return err
	}
	db := database.GetInstance()
	for _, s := range skills {
		history := &PlayerHistory{
			PlayerID: playerID,
			Skill:    s.Skill,
			Value:    s.Value,
		}

		// find skill combination
		var find PlayerHistory
		err = db.Where(&PlayerHistory{PlayerID: playerID, Skill: s.Skill, Value: s.Value}).First(&find).Error
		if err != nil && err != gorm.ErrRecordNotFound {
			return err
		}
		// if not found the same combination
		if find.ID < 1 {
			err = db.Save(history).Error
			if err != nil {
				defer db.Close()
				return err
			}
		}
	}
	defer db.Close()
	return nil
}

func deleteSkills(playerID uint) error {

	//@TODO colocar no belongsToCareerAndAcoount

	db := database.GetInstance()
	skills, err := GetSkills(playerID)
	if err != nil {
		return err
	}

	for _, skill := range skills {
		err = db.Delete(skill).Error
		if err != nil {
			defer db.Close()
			return err
		}
	}
	defer db.Close()
	return nil
}

func (player *Player) Update() (*Player, error) {

	if err := player.belongsToCareerAndAcoount(); err != nil {
		return nil, err
	}

	if err := player.Validate(); err != nil {
		return nil, err
	}

	// remove skills
	if err := deleteSkills(player.ID); err != nil {
		return nil, err
	}

	//update player
	db := database.GetInstance()
	err := db.Save(player).Error
	defer db.Close()
	if err != nil {
		return nil, err
	}

	// Save History
	if err := saveHistory(player.ID); err != nil {
		return nil, err
	}

	return player, err
}

func GetByID(id uint) (*Player, error) {

	//@TODO colocar no belongsToCareerAndAcoount

	player := &Player{}
	db := database.GetInstance()

	err := db.First(&player, id).Error
	if err != nil {
		defer db.Close()
		return nil, err
	}

	if player.Skills, err = GetSkills(player.ID); err != nil {
		return nil, err
	}

	defer db.Close()
	return player, nil
}

func (player *Player) GetSkills2() error {

	//@TODO colocar no belongsToCareerAndAcoount
	//@TODO ver por que tem 2

	skills, err := GetSkills(player.ID)
	if err != nil {
		return err
	}
	player.Skills = skills
	return nil
}

func GetSkills(playerId uint) ([]PlayerSkill, error) {

	//@TODO colocar no belongsToCareerAndAcoount
	var skills []PlayerSkill

	db := database.GetInstance()
	err := db.Where("player_id = ?", playerId).Find(&skills).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		defer db.Close()
		return nil, err
	}

	return skills, nil
}

func (player *Player) Validate() (err error) {

	var minPoints = 1
	var maxPoints = 100

	var minAge = 14
	var maxAge = 99

	if err = player.belongsToCareerAndAcoount(); err != nil {
		return err
	}

	if player.ID < 1 {
		if err = validatePlayerCombination(player.CareerID, player.Name, player.Position); err != nil {
			return err
		}
	}

	if strings.TrimSpace(player.Name) == "" {
		err = errors.New("name is required")
		log.Error(err)
		return err
	}

	if _, ok := Positions[player.Position]; !ok {
		err = errors.New("position is not valid")
		log.Error(err)
		return err
	}

	if _, ok := StrongFoots[player.StrongFoot]; !ok {
		err = errors.New("strong foot is not valid")
		log.Error(err)
		return err
	}

	if player.CareerID < 1 {
		err = errors.New("career is required")
		log.Error(err)
		return err
	}

	if player.PointsInitial < minPoints && player.PointsInitial > maxPoints {
		err = errors.New("initial points needs 1 between 100 ")
		log.Error(err)
		return err
	}

	if player.PotentialInitial < minPoints && player.PotentialInitial > maxPoints {
		err = errors.New("potential initial points needs 1 between 100")
		log.Error(err)
		return err
	}

	if player.PotentialFinal < minPoints && player.PotentialFinal > maxPoints {
		err = errors.New("potential final points needs 1 between 100")
		log.Error(err)
		return err
	}

	if player.AgeInitial < minAge && player.AgeInitial > maxAge {
		err = errors.New("age initial needs 14 between 99")
		log.Error(err)
		return err
	}

	if player.AgeCurrent < minAge && player.AgeCurrent > maxAge {
		err = errors.New("age current needs 14 between 99")
		log.Error(err)
		return err
	}

	for _, skill := range player.Skills {
		if _, ok := Skills[skill.Skill]; !ok {
			err = errors.New("skill is not valid: " + skill.Skill)
			log.Error(err)
			return err
		}
	}

	return err
}

func validatePlayerCombination(careerID uint, name string, position string) error {
	var find Player

	db := database.GetInstance()
	db.Where(&Player{Name: name, CareerID: careerID, Position: position}).First(&find)
	defer db.Close()

	if find.ID > 0 {
		return errors.New("player combination already exists")
	}

	return nil
}

func (player *Player) Delete() error {

	_, err := utils.GetCurrentUser()
	if err != nil {
		return err
	}

	db := database.GetInstance()
	var playerFind Player
	err = db.Find(&playerFind, "id = ? ", player.ID).Error
	defer db.Close()
	if err != nil {
		return err
	}

	if err = player.belongsToCareerAndAcoount(); err != nil {
		return err
	}

	db = database.GetInstance()
	err = db.Delete(&player).Error
	defer db.Close()
	if err != nil {
		return err
	}

	return err
}

func (player *Player) List() (players []Player, err error) {
	db := database.GetInstance()

	err = db.Order("name").Find(&players).Error
	if err != nil {
		defer db.Close()
		return players, err
	}

	defer db.Close()
	return players, err
}
