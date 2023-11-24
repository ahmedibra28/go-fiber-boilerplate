package models

import (
	"errors"

	"gorm.io/gorm"
)

type Permission struct {
	gorm.Model
	Name        string `json:"name" gorm:"not null"`
	Method      string `json:"method" gorm:"not null"`
	Route       string `json:"route" gorm:"not null"`
	Description string `json:"description"`

	Roles []Role `json:"roles" gorm:"many2many:role_permissions"`
}

func (p *Permission) BeforeSave(tx *gorm.DB) error {
	var existing Permission
	err := tx.Where("method = ? AND route = ? AND id != ?", p.Method, p.Route, p.ID).First(&existing).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil
	}

	return err
}
