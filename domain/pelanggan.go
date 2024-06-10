package domain

import (
	"time"

	"gorm.io/gorm"
)

type Pelanggan struct {
	ID              uint           `gorm:"primarykey" json:"id"`
	UserId          uint           `gorm:"not null" json:"user_id"`
	NoPelanggan     string         `gorm:"not null" json:"no_pelanggan"`
	NamaPelanggan   string         `gorm:"not null" json:"nama_pelanggan"`
	NikPelanggan    string         `gorm:"not null" json:"nik_pelanggan"`
	AlamatPelanggan string         `gorm:"not null" json:"alamat_pelanggan"`
	KecamatanId     *uint          `gorm:"null" json:"kecamatan_id"`
	KelurahanId     *uint          `gorm:"null" json:"kelurahan_id"`
	GolonganId      uint           `gorm:"not null" json:"golongan_id"`
	CreatedAt       *time.Time     `json:"created_at"`
	UpdatedAt       *time.Time     `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

type InsertData struct {
	UserID          uint   `gorm:"not null" json:"user_id"`
	PelangganID     uint   `gorm:"not null" json:"pelanggan_id"`
	NamaPelanggan   string `gorm:"not null" json:"nama_pelanggan"`
	AlamatPelanggan string `gorm:"not null" json:"alamat_pelanggan"`
}

type PelangganRepository interface {
	CreatePelanggan(req *Pelanggan) (*Pelanggan, error)
}

type PelangganUseCase interface {
	CreateImportPelanggan(req *InsertData) error
}
