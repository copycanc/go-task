package storages

import (
	"go-br-task/internal/models"

	"github.com/google/uuid"
)

type MapStorageUser struct {
	storageUserMap map[uuid.UUID]models.User
}

func NewMapStorageUser() *MapStorageUser {
	return &MapStorageUser{storageUserMap: make(map[uuid.UUID]models.User)}
}
func (m *MapStorageUser) GetAllUser() (map[uuid.UUID]models.User, error) {
	return m.storageUserMap, nil
}

func (m *MapStorageUser) ExistEmailUser(email string) (bool, error) {
	for _, user := range m.storageUserMap {
		if user.Email == email {
			return true, nil
		}
	}
	return false, nil
}

func (m *MapStorageUser) SaveUser(user models.User) error {
	m.storageUserMap[user.ID] = user
	return nil
}

func (m *MapStorageUser) GetUserID(id uuid.UUID) (*models.User, error) {
	user, _ := m.storageUserMap[id]
	return &user, nil
}

func (m *MapStorageUser) ExistUser(id uuid.UUID) (bool, error) {
	_, exist := m.storageUserMap[id]
	if !exist {
		return false, nil
	}
	return true, nil
}

func (m *MapStorageUser) DeleteUser(id uuid.UUID) error {
	delete(m.storageUserMap, id)
	return nil
}
