package models

import (
	"gorm.io/gorm"
)

type RolePermission struct {
	gorm.Model

	RoleID       uint `json:"role_id"`
	PermissionID uint `json:"permission_id"`
}
