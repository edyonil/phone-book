package service

import (
	"phoneService/config"
	"phoneService/model"
)

type PhoneLister interface {
	GetAll(term string) ([]model.Phone, error)
}

type phoneListService struct{}

func New() PhoneLister {
	return &phoneListService{}
}

func (p *phoneListService) GetAll(term string) ([]model.Phone, error) {

	var db = config.Db.DB

	var phones []model.Phone

	rows := db.
		Preload("Employee").
		Joins("LEFT JOIN employees ON employees.id = phones.employee_id").
		Joins("LEFT JOIN units as unitE ON unitE.id = employees.unit_id").
		Joins("LEFT JOIN units ON units.id = phones.unit_id").
		Where("employees.name like ?", "%"+term+"%").
		Or("unitE.name like ?", "%"+term+"%").
		Preload("Employee.Unit").
		Debug().
		Find(&phones)

	if rows.Error != nil {
		return phones, rows.Error
	}

	return phones, nil
}
