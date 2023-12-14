package models

import (
	"gorm.io/gorm"
)

type Basket struct {
	gorm.Model
	UserID uint   `gorm:"not null"` // Link to the User model
	Data   string `gorm:"type:varchar(2048)" validate:"max=2048"`
	State  string `gorm:"type:varchar(10)" validate:"customState"`
}

type BasketInput struct {
	Data  string
	State string
}
