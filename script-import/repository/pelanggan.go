package repository

import (
	"script-import/domain"

	"gorm.io/gorm"
)

type mysqlPelanggan struct {
	DB *gorm.DB
}

func NewPostgrePelanggan(client *gorm.DB) domain.PelangganRepository {
	return &mysqlPelanggan{
		DB: client,
	}
}

func (p *mysqlPelanggan) CreatePelanggan(req *domain.Pelanggan) (*domain.Pelanggan, error) {
	err := p.DB.
		Create(&req).
		Error

	if err != nil {
		return &domain.Pelanggan{}, err
	}

	createdPelanggan := &domain.Pelanggan{}
	err = p.DB.
		Last(createdPelanggan).
		Error

	if err != nil {
		return &domain.Pelanggan{}, err
	}

	return createdPelanggan, nil
}
