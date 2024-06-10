package repository

import (
	"script-import/domain"

	"gorm.io/gorm"
)

type mysqlUser struct {
	DB *gorm.DB
}

func NewPostgreUser(client *gorm.DB) domain.UserRepository {
	return &mysqlUser{
		DB: client,
	}
}

func (p *mysqlUser) CreateUser(req *domain.User) (*domain.User, error) {
	err := p.DB.
		Create(&req).
		Error

	if err != nil {
		return &domain.User{}, err
	}

	createdUser := &domain.User{}
	err = p.DB.
		Last(createdUser).
		Error

	if err != nil {
		return &domain.User{}, err
	}

	return createdUser, nil
}
