package models

import (
	"final-project-gin/helpers"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	GormModel
	FullName string `gorm:not null" json:"full_name" form:"full_name" valid: required-your full name is required"`
	Email    string `gorm:not null;uniqueIndex" json:"email" form:"email" valid: required-your email is required,email-Invalid email format"`
	Password string `gorm:not null json:"password" form:"password" valid: required-your password is required,minstringlegth(6)~password has to have a minimum length of 6 characters"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}
	u.Password = helpers.HashPass(u.Password)
	err = nil
	return
}
