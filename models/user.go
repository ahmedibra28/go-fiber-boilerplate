package models

import (
	"github.com/ahmedibra28/go-fiber-boilerplate/utils"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name                string `json:"name" gorm:"not null"`
	Email               string `json:"email" gorm:"unique;not null"`
	Image               string `json:"image"`
	Mobile              string `json:"mobile"`
	Address             string `json:"address"`
	Bio                 string `json:"bio"`
	Password            string `json:"password" gorm:"not null"`
	Confirmed           bool   `json:"confirmed" default:"false"`
	Blocked             bool   `json:"blocked" default:"false"`
	ResetPasswordToken  string `json:"reset_password_token"`
	ResetPasswordExpire int64  `json:"reset_password_expire"`

	Role   Role `json:"role" gorm:"foreignKey:RoleID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	RoleID uint `json:"role_id"`
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
