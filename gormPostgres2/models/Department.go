package models

import "gorm.io/gorm"

type Department struct {
	gorm.Model
	DepartmentName string `json:"DepartmentName"`
}
