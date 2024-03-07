package models

import (
	"gorm.io/gorm"
)

type Car struct {
	gorm.Model
	Name       string
	Make       string
	ModelYear  int 
	TypeID     int
	Type       CarType
	Color      string
	SpeedRange string
}
