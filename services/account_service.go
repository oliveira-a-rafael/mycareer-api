package services

import (
	"fmt"

	"github.com/oliveira-a-rafael/mycareer-api/domains"
	"github.com/oliveira-a-rafael/mycareer-api/infrastructure/repositories"
)

type AccountContract interface {
	CreateAccount(account *domains.Account) error
}

type AccountService struct {
	AccountRepository repositories.AccountQuerier `inject:""`
}

func (s *AccountService) CreateAccount(account *domains.Account) error {
	fmt.Println("account: ", account)
	// return s.AccountRepository.Insert(account)
	return nil
}
