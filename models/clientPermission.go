package models

import "time"

type ClientPermission struct {
	ID        uint      `json:"id,omitempty" gorm:"primaryKey;autoIncrement"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`

	Name        string `json:"name,omitempty" gorm:"not null;default:null"`
	Sort        int    `json:"sort,omitempty" gorm:"not null;default:null"`
	Menu        string `json:"menu,omitempty" gorm:"not null;default:null"`
	Path        string `json:"path,omitempty" gorm:"not null;unique;default:null"`
	Description string `json:"description,omitempty"`

	Roles []Role `json:"roles,omitempty" gorm:"many2many:role_client_permissions"`
}
