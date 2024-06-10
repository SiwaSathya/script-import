package usecase

import (
	"fmt"
	"script-import/domain"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type pelangganUseCase struct {
	pelangganRepository domain.PelangganRepository
	userRepository      domain.UserRepository
	contextTimeout      time.Duration
}

func NewPelangganUseCase(dtl domain.PelangganRepository, usr domain.UserRepository, t time.Duration) domain.PelangganUseCase {
	return &pelangganUseCase{
		pelangganRepository: dtl,
		userRepository:      usr,
		contextTimeout:      t,
	}
}

func (p *pelangganUseCase) CreateImportPelanggan(req *domain.InsertData) error {

	trimmedString := strings.ReplaceAll(req.NamaPelanggan, " ", "")
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(strings.ToLower(trimmedString)), bcrypt.DefaultCost)
	if err != nil {
		// fmt.Println(err)
		return err
	}

	usr := domain.User{
		ID:              req.UserID,
		Name:            req.NamaPelanggan,
		Email:           fmt.Sprintf("%s%d@gmail.com", strings.ToLower(trimmedString), req.UserID),
		Password:        string(hashedPassword),
		EmailVerifiedAt: nil,
		DeletedAt:       nil,
	}
	res, err := p.userRepository.CreateUser(&usr)
	if err != nil {
		return err
	}
	pel := domain.Pelanggan{
		ID:              req.PelangganID,
		UserId:          res.ID,
		NoPelanggan:     "",
		NamaPelanggan:   req.NamaPelanggan,
		AlamatPelanggan: req.AlamatPelanggan,
		KecamatanId:     nil,
		KelurahanId:     nil,
		GolonganId:      2,
	}
	_, err = p.pelangganRepository.CreatePelanggan(&pel)
	if err != nil {
		return err
	}

	return nil
}
