package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Title       string `gorm:"not null" json:"title" form:"title" valid:"required~Title of your product is required"`
	Description string `gorm:"not null" json:"description" form:"description" valid:"required~Description of your product is required"`
	UserID      uint
	User        *User
}

func (p *Product) BeforeCreate(tx *gorm.DB) error {
	_, err := govalidator.ValidateStruct(p)
	if err != nil {
		return err
	}

	return nil
}

// func (p *Product) BeforeUpdate(tx *gorm.DB) error {
// 	_, err := govalidator.ValidateStruct(p)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
