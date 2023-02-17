package model

import "github.com/google/uuid"

type CompanyInfo struct {
	Id               uuid.UUID `gorm:"PrimaryKey" json:"id"`
	CompanyName      string    `gorm:"column:name" json:"name" validate:"required"`
	Description      string    `gorm:"column:description" json:"description"`
	EmployeeStrength int       `gorm:"column:employee_count" json:"employee_strength" validate:"required"`
	Registered       bool      `gorm:"column:registered" json:"registered" validate:"required"`
	CompanyType      string    `gorm:"column:type" json:"type" validate:"required"`
}
