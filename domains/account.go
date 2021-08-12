package domains

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
)

type Token struct {
	UserId uint
	jwt.StandardClaims
}

type Account struct {
	gorm.Model
	Document string   `gorm:"unique_index;not null" json:"document"`
	Name     string   `json:"name"`
	Email    string   `gorm:"type:varchar(100); unique_index" json:"email"`
	Password string   `json:"password"`
	Token    string   `json:"token"`
	Career   []Career `gorm:"foreignkey:AccountID"`
}
