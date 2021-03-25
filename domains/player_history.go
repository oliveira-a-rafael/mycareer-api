package domains

import (
	"github.com/jinzhu/gorm"
)

// PlayerHistory model
type PlayerHistory struct {
	gorm.Model
	Player   Player `gorm:"foreignkey:PlayerID"`
	PlayerID uint   `gorm:"index:player_history_idx"`
	Skill    string `gorm:"index:skill_idx"`
	Value    int
}
