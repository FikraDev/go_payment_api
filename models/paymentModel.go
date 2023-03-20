package models

import (
	"gorm.io/gorm"
)

type Payment struct {
	gorm.Model
	Amount      float64
	Quantity    float64
	Description string
	ItemCode    string
}
