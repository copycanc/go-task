package user

import (
	"github.com/google/uuid"
)

type UserStorage interface {
	GetAllUser() (map[uuid.UUID]User, error)
	ExistEmailUser(email string) (bool, error)
	SaveUser(user User) error
	GetUserID(id uuid.UUID) (*User, error)
	ExistUser(id uuid.UUID) (bool, error)
	DeleteUser(id uuid.UUID) error
	GetUserName(name string) (*User, error)
}
