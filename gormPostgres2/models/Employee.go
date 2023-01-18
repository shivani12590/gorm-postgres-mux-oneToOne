package models

import "gorm.io/gorm"

type Employee struct {
	gorm.Model
	FirstName    string     `json:"firstName"`
	LastName     string     `json:"lastName"`
	Department   Department `json:"department"`
	DepartmentId uint       `gorm:"foreignKey:ID" json:"_"`
}
