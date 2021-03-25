package services

import (
	"github.com/oliveira-a-rafael/mycareer-api/domains"
	"github.com/oliveira-a-rafael/mycareer-api/infrastructure/repositories"
)

type AccountContract interface {
	CreateAccount(account *domains.Account) error
}

type AccountService struct {
	AccountRepository repositories.AccountQuerier
}

func (s *AccountService) CreateAccount(account *domains.Account) error {
	return s.AccountRepository.Insert(account)
}
