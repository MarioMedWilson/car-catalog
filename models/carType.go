package models

import (
	"gorm.io/gorm"
)

type CarType struct {
	gorm.Model
	Type	string	`json:"type"`
}
