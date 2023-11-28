package models

import (
	"time"

	"github.com/ahmedibra28/go-fiber-boilerplate/utils"
	"gorm.io/gorm"
)

type User struct {
	ID        uint      `json:"id,omitempty" gorm:"primaryKey;autoIncrement"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`

	Name                string `json:"name,omitempty" gorm:"not null;default:null"`
	Email               string `json:"email,omitempty" gorm:"unique;not null;default:null"`
	Image               string `json:"image,omitempty"`
	Mobile              string `json:"mobile,omitempty"`
	Address             string `json:"address,omitempty"`
	Bio                 string `json:"bio,omitempty"`
	Password            string `json:"password,omitempty" gorm:"not null;default:null"`
	Confirmed           bool   `json:"confirmed,omitempty" default:"false"`
	Blocked             bool   `json:"blocked,omitempty" default:"false"`
	ResetPasswordToken  string `json:"reset_password_token,omitempty"`
	ResetPasswordExpire int64  `json:"reset_password_expire,omitempty"`

	Role   Role `json:"role,omitempty" gorm:"foreignKey:RoleID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	RoleID uint `json:"role_id,omitempty"`
}

func (u *User) BeforeSave(db *gorm.DB) error {
	if u.Password != "" {
		hashedPassword, err := utils.HashPassword(u.Password)
		if err != nil {
			return err
		}
		u.Password = hashedPassword
	}

	return nil
}
