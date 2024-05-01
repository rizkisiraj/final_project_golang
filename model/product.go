package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	UUID     string    `gorm:"not null" json:"uuid"`
	Name     string    `gorm:"not null" json:"name" form:"name" valid:"required~Product name is required"`
	ImageURL string    `json:"image_url" valid:"required~Image url is required"`
	AdminID  uint      `gorm:"not null" json:"admin_id" valid:"required~Admin id is required"`
	Variants []Variant `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"variants"`
}
