package entity

import (
	"adamnasrudin03/challenge-wallet/pkg/helpers"
	"errors"
	"time"

	"github.com/go-playground/validator/v10"

	"gorm.io/gorm"
)

type User struct {
	ID        uint64    `gorm:"primaryKey" json:"id"`
	FullName  string    `gorm:"not null" json:"full_name" validate:"required,min=3"`
	Email     string    `gorm:"not null;uniqueIndex" json:"email" validate:"required,email"`
	Password  string    `gorm:"not null" json:"password,omitempty" validate:"required,min=6"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	validate := validator.New()
	err = validate.Struct(u)
	if err != nil {
		errs := helpers.FormatValidationError(err)
		return errors.New(errs)
	}

	hashedPass, err := helpers.HashPassword(u.Password)
	if err != nil {
		return
	}
	u.Password = hashedPass

	return
}
