package auth

import "go-br-task/internal/user"

type AuthService struct {
	storage user.UserStorage
}

func NewAuthService(storage user.UserStorage) *AuthService {
	return &AuthService{
		storage: storage,
	}
}
