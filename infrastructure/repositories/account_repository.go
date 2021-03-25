package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/oliveira-a-rafael/mycareer-api/domains"
)

type AccountQuerier interface {
	Insert(model *domains.Account) error
	// FindAllByAttendantIdAndStoreIdAndStatus(list *[]*domains.Attendance, attendantID string, storeID uint, status string) error
	// GetById(id string, model *domains.Attendance) error
	// GetByAccountId(accountID uint, model *domains.Attendance) error
	// GetByCurrentAttendance(currentAttendance uint, model *domains.Attendance) error
	// Update(model *domains.Attendance) error
	// Delete(model *domains.Attendance) error
	// DeleteById(id string) error
	// CountByCustomer(attendantID string, attendantCustomerID string, status string) (error, int)
}

type AccountRepository struct {
	DB *gorm.DB
}

func (r *AccountRepository) Insert(model *domains.Career) error {
	err := r.DB.Create(model).Error
	// @TODO study application monitoring platform
	//return sentry.HandleError(err)
	return err
}
