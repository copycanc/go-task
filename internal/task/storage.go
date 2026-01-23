package task

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type MapStorageTask struct {
	storageTaskMap map[uuid.UUID]Task
}

type PostgresqlStorageTask struct {
	pool *pgxpool.Pool
}

func NewPGStorageTask(pool *pgxpool.Pool) *PostgresqlStorageTask {
	return &PostgresqlStorageTask{pool: pool}
}

func NewMapStorageTask() *MapStorageTask {
	return &MapStorageTask{storageTaskMap: make(map[uuid.UUID]Task)}
}

func (m *MapStorageTask) GetAllTask() (map[uuid.UUID]Task, error) {
	return m.storageTaskMap, nil
}

func (p *PostgresqlStorageTask) GetAllTask() (map[uuid.UUID]Task, error) {
	ctx := context.Background()
	sqlReq := `SELECT id, title, description, status, created_at, completed_at FROM tasks`
	rows, err := p.pool.Query(ctx, sqlReq)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tasks := make(map[uuid.UUID]Task)

	for rows.Next() {
		var t Task
		if err := rows.Scan(
			&t.ID,
			&t.Title,
			&t.Description,
			&t.Status,
			&t.CreatedAt,
			&t.CompletedAt,
		); err != nil {
			return nil, err
		}
		tasks[t.ID] = t
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (m *MapStorageTask) SaveTask(task Task) error {
	m.storageTaskMap[task.ID] = task
	return nil
}

func (p *PostgresqlStorageTask) SaveTask(task Task) error {
	ctx := context.Background()
	sqlReq := `INSERT INTO tasks(
                  id,title,description,status,created_at,completed_at) VALUES ($1,$2,$3,$4,$5,$6) ON CONFLICT (id) DO UPDATE SET
					title = EXCLUDED.title,
					description = EXCLUDED.description,
					status = EXCLUDED.status,
					completed_at = EXCLUDED.completed_at`
	_, err := p.pool.Exec(ctx, sqlReq, task.ID, task.Title, task.Description, task.Status, task.CreatedAt, task.CompletedAt)
	if err != nil {
		return err
	}
	return nil
}

func (m *MapStorageTask) ExistTask(uuid uuid.UUID) (bool, error) {
	_, exist := m.storageTaskMap[uuid]
	if !exist {
		return false, nil
	}
	return true, nil
}
func (p *PostgresqlStorageTask) ExistTask(uuid uuid.UUID) (bool, error) {
	ctx := context.Background()
	var exist bool
	err := p.pool.QueryRow(ctx, `SELECT exists(SELECT 1 FROM tasks WHERE id=$1 )`, uuid).Scan(&exist)
	if err != nil {
		return false, err
	}
	return exist, nil
}

func (m *MapStorageTask) GetTaskID(uuid uuid.UUID) (*Task, error) {
	task, _ := m.storageTaskMap[uuid]
	return &task, nil
}

func (p *PostgresqlStorageTask) GetTaskID(uuid uuid.UUID) (*Task, error) {
	var task Task
	ctx := context.Background()
	sqlReq := `SELECT id, title, description, status, created_at, completed_at FROM tasks WHERE id = $1`
	err := p.pool.QueryRow(ctx, sqlReq, uuid).Scan(&task.ID, &task.Title,
		&task.Description,
		&task.Status,
		&task.CreatedAt,
		&task.CompletedAt)
	if err != nil {
		return nil, err
	}
	return &task, nil
}

func (m *MapStorageTask) DeleteTask(uuid uuid.UUID) error {
	delete(m.storageTaskMap, uuid)
	return nil
}

func (p *PostgresqlStorageTask) DeleteTask(uuid uuid.UUID) error {
	ctx := context.Background()
	_, err := p.pool.Exec(ctx, `DELETE FROM tasks WHERE id = $1`, uuid)
	if err != nil {
		return err
	}
	return nil
}
