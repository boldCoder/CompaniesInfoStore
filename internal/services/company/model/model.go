package model

import "github.com/google/uuid"

type CompanyInfo struct {
	Id               uuid.UUID `gorm:"primary_key" json:"id"`
	CompanyName      string    `gorm:"column:name;not null" json:"name" validate:"required"`
	Description      string    `gorm:"column:description" json:"description"`
	EmployeeStrength int       `gorm:"column:employee_count;not null" json:"employee_strength" validate:"required"`
	Registered       bool      `gorm:"column:registered;not null" json:"registered" validate:"required"`
	CompanyType      string    `gorm:"column:type;not null" json:"type" validate:"required"`
}
