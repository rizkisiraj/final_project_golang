package model

import "gorm.io/gorm"

type Admin struct {
	gorm.Model
	UUID     string    `gorm:"not null" json:uuid`
	Name     string    `gorm:"not null" json:name form:"name" valid:"required~Your full name is required"`
	Email    string    `gorm:"not null" json:email form:"email" valid:"required~Your email is required, email~Invalid email format"`
	Password string    `gorm:"not null" json:"password" form:"password" valid:"required~Your password is required, minstringlength(6)~Password has to have a minimum length of 6 characters"`
	Products []Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"products"`
}
