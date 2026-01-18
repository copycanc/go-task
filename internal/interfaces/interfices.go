package interfaces

import (
	"go-br-task/internal/models"

	"github.com/google/uuid"
)

type TasksStorage interface {
	GetAllTask() (map[uuid.UUID]models.Task, error)
	SaveTask(models.Task) error
	GetTaskID(id uuid.UUID) (*models.Task, error)
	ExistTask(id uuid.UUID) (bool, error)
	DeleteTask(id uuid.UUID) error
}

type UserStorage interface {
	GetAllUser() (map[uuid.UUID]models.User, error)
	ExistEmailUser(email string) (bool, error)
	SaveUser(user models.User) error
	GetUserID(id uuid.UUID) (*models.User, error)
	ExistUser(id uuid.UUID) (bool, error)
	DeleteUser(id uuid.UUID) error
}
