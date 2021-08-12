package repositories

import (
	"errors"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"github.com/oliveira-a-rafael/mycareer-api/database"
	"github.com/oliveira-a-rafael/mycareer-api/domains"
	"github.com/oliveira-a-rafael/mycareer-api/utils"
)

type AccountRepository struct {
	DB *gorm.DB
}

func (r *AccountRepository) GetById(id int, model *domains.Account) error {
	err := r.DB.Where("id = ?", id).Find(&model).Error
	return err
}

func (r *AccountRepository) Create(account *domains.Account) (*domains.Account, error) {
	err := r.validateModel(account)
	if err != nil {
		return account, err
	}

	hashedPassword := utils.BuildHashedPassword(account.Password)
	account.Password = string(hashedPassword)

	r.DB.Create(account)
	if err != nil {
		//@TODO logar os erros nessa camada
		return account, err
	}

	//Create new JWT token for the newly registered account
	tk := &domains.Token{UserId: account.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	account.Token = tokenString

	account.Password = "" //delete password

	return account, err
}

func (r *AccountRepository) validateModel(account *domains.Account) (err error) {
	// @TODO colocar validate ver utro projeto
	if !strings.Contains(account.Email, "@") {
		err = errors.New("e-mail address is // required")
		//log.Error(err)
		return err
	}

	if len(account.Password) < 6 {
		err = errors.New("password is required")
		// log.Error(err)
		return err
	}

	// Validate email
	temp := &domains.Account{}
	db := database.GetInstance()
	err = db.Table("accounts").Where("email = ?", account.Email).First(temp).Error
	db.Close()
	if err != nil && err != gorm.ErrRecordNotFound {
		// log.Error(err)
		return err
	}
	if temp.Email != "" {
		err = errors.New("e-mail address already in use by another user")
		// log.Error(err)
		return err
	}
	err = nil

	// Validate document
	temp = &domains.Account{}
	db = database.GetInstance()
	err = db.Table("accounts").Where("document = ?", account.Document).First(temp).Error
	db.Close()
	if err != nil && err != gorm.ErrRecordNotFound {
		// log.Error(err)
		return err
	}
	if temp.Document != "" {
		err = errors.New("document already exists")
		// log.Error(err)
		return err
	}
	err = nil

	return err
}
