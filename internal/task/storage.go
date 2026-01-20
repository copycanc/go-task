package task

import (
	"github.com/google/uuid"
)

type MapStorageTask struct {
	storageTaskMap map[uuid.UUID]Task
}

func NewMapStorageTask() *MapStorageTask {
	return &MapStorageTask{storageTaskMap: make(map[uuid.UUID]Task)}
}

func (m *MapStorageTask) GetAllTask() (map[uuid.UUID]Task, error) {
	return m.storageTaskMap, nil
}

func (m *MapStorageTask) SaveTask(task Task) error {
	m.storageTaskMap[task.ID] = task
	return nil
}

func (m *MapStorageTask) ExistTask(uuid uuid.UUID) (bool, error) {
	_, exist := m.storageTaskMap[uuid]
	if !exist {
		return false, nil
	}
	return true, nil
}

func (m *MapStorageTask) GetTaskID(uuid uuid.UUID) (*Task, error) {
	task, _ := m.storageTaskMap[uuid]
	return &task, nil
}

func (m *MapStorageTask) DeleteTask(uuid uuid.UUID) error {
	delete(m.storageTaskMap, uuid)
	return nil
}
