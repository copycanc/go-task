package utils

import (
	"go-br-task/internal/models"
)

func ChekChangePass(u models.ChangeUser) bool {
	if u.NewPassword != "" && u.OldPassword != "" {
		return true
	}
	return false
}

func ChekChangeEmail(u models.ChangeUser) bool {
	if u.Email != "" {
		return true
	}
	return false
}
