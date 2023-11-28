package models

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Permission struct {
	ID        uint      `json:"id,omitempty" gorm:"primaryKey;autoIncrement"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`

	Name        string `json:"name,omitempty" gorm:"not null;default:null"`
	Method      string `json:"method,omitempty" gorm:"not null;default:null"`
	Route       string `json:"route,omitempty" gorm:"not null;default:null"`
	Description string `json:"description,omitempty"`

	Roles []Role `json:"roles,omitempty" gorm:"many2many:role_permissions"`
}

func (p *Permission) BeforeSave(tx *gorm.DB) error {
	var count int64
	result := tx.Model(&Permission{}).Where("method = ? AND route = ? AND id <> ?", p.Method, p.Route, p.ID).Count(&count)
	if result.Error != nil {
		return result.Error
	}

	if count > 0 {
		return fmt.Errorf("method with route already exists")
	}

	return nil
}
