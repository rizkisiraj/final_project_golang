package model

import "gorm.io/gorm"

type Variant struct {
	gorm.Model
	UUID        string `gorm:"not null" json:"uuid"`
	VariantName string `gorm:"not null" json:"name" form:"name" valid:"required~Product name is required"`
	Quantity    int    `gorm:"not null" json:"quantity" form:"quantity" valid:"required~Variant quantity is required"`
	ProductId   int    `gorm:"not null" json:"product_id" form:"product_id" valid:"required~Product ID is required"`
}
