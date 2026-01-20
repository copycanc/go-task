package user

import "github.com/google/uuid"

// Структура пользователей
type User struct {
	ID       uuid.UUID
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Структура для изменения пользователей
type ChangeUser struct {
	Email       string `json:"email"`
	NewPassword string `json:"newpassword"`
	OldPassword string `json:"oldpassword"`
}

// Структура для показа пользователей
type UserOutput struct {
	ID    uuid.UUID
	Name  string
	Email string
}

// Преобразовываем пользователя, что бы показывать без пароля
func (u *User) OutputUser() UserOutput {
	return UserOutput{
		ID:    u.ID,
		Name:  u.Name,
		Email: u.Email,
	}
}
