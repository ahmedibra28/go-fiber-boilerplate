package utils

import (
	"github.com/gofiber/fiber/v2"
)

type UserInfo struct {
	User User `json:"user"`
	Role Role `json:"role"`
}

type User struct {
	ID     uint   `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Mobile string `json:"mobile"`
}

type Role struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

func GetUserInfo(c *fiber.Ctx) UserInfo {

	userID := c.Locals("userID").(uint)
	userName := c.Locals("userName").(string)
	userMobile := c.Locals("userMobile").(string)
	userEmail := c.Locals("userEmail").(string)
	roleID := c.Locals("roleID").(uint)
	roleName := c.Locals("roleName").(string)
	roleType := c.Locals("roleType").(string)

	return UserInfo{
		User: User{
			ID:     userID,
			Name:   userName,
			Email:  userEmail,
			Mobile: userMobile,
		},
		Role: Role{
			ID:   roleID,
			Name: roleName,
			Type: roleType,
		},
	}
}
