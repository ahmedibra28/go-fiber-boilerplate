package models

import (
	"strings"
	"time"

	"gorm.io/gorm"
)

type Role struct {
	ID        uint      `json:"id,omitempty" gorm:"primaryKey;autoIncrement"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`

	Name        string `json:"name,omitempty" gorm:"not null;unique;default:null"`
	Type        string `json:"type,omitempty" gorm:"not null;unique;default:null"`
	Description string `json:"description,omitempty"`

	Users             []User             `json:"users,omitempty" gorm:"foreignKey:RoleID"`
	Permissions       []Permission       `json:"permissions,omitempty" gorm:"many2many:role_permissions"`
	ClientPermissions []ClientPermission `json:"client_permissions,omitempty" gorm:"many2many:role_client_permissions"`
}

func (r *Role) BeforeSave(tx *gorm.DB) (err error) {

	var roleType = strings.ReplaceAll(r.Name, " ", "_")

	r.Type = strings.ToUpper(roleType)

	return nil
}
