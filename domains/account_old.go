package domains

import (
	"errors"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"github.com/oliveira-a-rafael/mycareer-api/database"
	"github.com/oliveira-a-rafael/mycareer-api/utils"
	"golang.org/x/crypto/bcrypt"
)

// type Token struct {
// 	UserId uint
// 	jwt.StandardClaims
// }

// User model
type AccountOld struct {
	gorm.Model
	Document string   `gorm:"unique_index;not null" json:"document"`
	Name     string   `json:"name"`
	Email    string   `gorm:"type:varchar(100); unique_index" json:"email"`
	Password string   `json:"password"`
	Token    string   `json:"token"`
	Career   []Career `gorm:"foreignkey:AccountID"`
}

func (account *AccountOld) Create() (*AccountOld, error) {
	err := account.Validate()
	if err != nil {
		return account, err
	}

	hashedPassword := utils.BuildHashedPassword(account.Password)
	account.Password = string(hashedPassword)

	db := database.GetInstance()
	db.Create(account)
	if err != nil {
		return account, err
	}

	//Create new JWT token for the newly registered account
	tk := &Token{UserId: account.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	account.Token = tokenString

	account.Password = "" //delete password

	defer db.Close()

	return account, err
}

func Login(email, password string) (*AccountOld, error) {

	account := &AccountOld{}
	db := database.GetInstance()
	err := db.Table("accounts").Where("email = ?", email).First(account).Error
	defer db.Close()
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return account, err
		}
		return account, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		err = errors.New("invalid login credentials. Please try again")
		return account, err
	}
	//Worked! Logged In
	account.Password = ""

	//Create JWT token
	tk := &Token{UserId: account.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	account.Token = tokenString //Store the token in the response

	return account, err
}

func (account *AccountOld) Validate() (err error) {

	if !strings.Contains(account.Email, "@") {
		err = errors.New("e-mail address is required")
		log.Error(err)
		return err
	}

	if len(account.Password) < 6 {
		err = errors.New("password is required")
		log.Error(err)
		return err
	}

	// Validate email
	temp := &AccountOld{}
	db := database.GetInstance()
	err = db.Table("accounts").Where("email = ?", account.Email).First(temp).Error
	db.Close()
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Error(err)
		return err
	}
	if temp.Email != "" {
		err = errors.New("e-mail address already in use by another user")
		log.Error(err)
		return err
	}
	err = nil

	// Validate document
	temp = &AccountOld{}
	db = database.GetInstance()
	err = db.Table("accounts").Where("document = ?", account.Document).First(temp).Error
	db.Close()
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Error(err)
		return err
	}
	if temp.Document != "" {
		err = errors.New("document already exists")
		log.Error(err)
		return err
	}
	err = nil

	return err
}

func (account *AccountOld) ListCareers() (err error, careers []Career) {

	db := database.GetInstance()
	err = db.Model(&account).Related(&careers).Error
	defer db.Close()

	for i, carrer := range careers {
		count := 0
		db.Model(&[]Player{}).Where(&Player{CareerID: carrer.ID}).Count(&count)
		careers[i].TotalPlayer = count
	}

	if err != nil {
		return err, careers
	}

	return err, careers
}

func TesteGetAccount() (*AccountOld, error) {

	account := &AccountOld{}
	db := database.GetInstance()
	id := 1
	err := db.Table("accounts").Where("id = ?", id).First(account).Error
	defer db.Close()
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return account, err
		}
		return account, err
	}

	return account, err
}
