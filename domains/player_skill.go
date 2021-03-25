package domains

import (
	"github.com/jinzhu/gorm"
)

// PlayerSkill model
type PlayerSkill struct {
	gorm.Model
	Player   Player `gorm:"foreignkey:PlayerID" json:"-"`
	PlayerID uint   `gorm:"index:player_skill_idx" json:"-"`
	Skill   string  `json:"skill"`
	Value    int    `json:"value"`
}