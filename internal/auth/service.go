package auth

import (
	"errors"
	"go-br-task/internal/user"
	"log/slog"

	"github.com/google/uuid"
)

type AuthService struct {
	storage user.UserStorage
}

func NewAuthService(storage user.UserStorage) *AuthService {
	return &AuthService{
		storage: storage,
	}
}

func (a *AuthService) CheckCredentials(email, password string) (uuid.UUID, int, error) {
	exist, err := a.storage.ExistEmailUser(email)
	if err != nil {
		slog.Error("STORAGE: get user failed", "err", err)
		return uuid.Nil, 500, errors.New("ошибка при получении данных")
	}
	if !exist {
		return uuid.Nil, 401, errors.New("Неверно введены логин или пароль")
	}
	users, errs := a.storage.GetUserEmail(email)
	if errs != nil {
		slog.Error("STORAGE: get user failed", "err", errs)
		return uuid.Nil, 500, errors.New("ошибка при получении данных")
	}
	if users.Password != password {
		return uuid.Nil, 401, errors.New("Неверно введены логин или пароль")
	}
	return users.ID, 200, nil
}
