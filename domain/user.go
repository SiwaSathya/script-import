package domain

import "time"

type User struct {
	ID              uint       `gorm:"primarykey" json:"id"`
	Name            string     `gorm:"not null" json:"name"`
	Email           string     `gorm:"not null" json:"email"`
	EmailVerifiedAt *time.Time `gorm:"null" json:"email_verified_at"`
	DeletedAt       *time.Time `gorm:"null" json:"deleted_at"`
	Password        string     `gorm:"not null" json:"password"`
	RememberToken   string     `gorm:"null" json:"remember_token"`
	CreatedAt       *time.Time `json:"created_at"`
	UpdatedAt       *time.Time `json:"updated_at"`
	DeviceToken     string     `gorm:"null" json:"device_token"`
}

type UserRepository interface {
	CreateUser(req *User) (*User, error)
}
