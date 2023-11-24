package models

import "gorm.io/gorm"

type ClientPermission struct {
	gorm.Model
	Name        string `json:"name" gorm:"not null"`
	Sort        int    `json:"sort" gorm:"not null"`
	Menu        string `json:"menu" gorm:"not null"`
	Path        string `json:"path" gorm:"not null;unique"`
	Description string `json:"description"`

	Roles []Role `json:"roles" gorm:"many2many:role_client_permissions"`
}
