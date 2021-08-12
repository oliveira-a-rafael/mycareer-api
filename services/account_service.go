package services

import (
	"github.com/oliveira-a-rafael/mycareer-api/domains"
	"github.com/oliveira-a-rafael/mycareer-api/infra/repositories"
)

type AccountContract interface {
	GetAccount()
	CreateAccount()
}

type AccountService struct {
	AccountRepository repositories.AccountRepository
}

func (as *AccountService) GetAccount(id int) (*domains.Account, error) {
	account := &domains.Account{}
	err := as.AccountRepository.GetById(id, account)
	return account, err
}

func (as *AccountService) CreateAccount(account *domains.Account) (*domains.Account, error) {
	return as.AccountRepository.Create(account)
}
