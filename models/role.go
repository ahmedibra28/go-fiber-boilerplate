package models

import (
	"strings"

	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	Name        string `json:"name" gorm:"not null;unique"`
	Type        string `json:"type" gorm:"not null;unique"`
	Description string `json:"description"`

	Users             []User             `json:"users" gorm:"foreignKey:RoleID"`
	Permissions       []Permission       `json:"permissions" gorm:"many2many:role_permissions"`
	ClientPermissions []ClientPermission `json:"client_permissions" gorm:"many2many:role_client_permissions"`
}

func (r *Role) BeforeSave(tx *gorm.DB) (err error) {
	r.Type = strings.ToUpper(r.Name)

	return nil
}
