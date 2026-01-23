package user

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type MapStorageUser struct {
	storageUserMap map[uuid.UUID]User
}

type PostgresqlStorageUser struct {
	pool *pgxpool.Pool
}

func NewMapStorageUser() *MapStorageUser {
	return &MapStorageUser{storageUserMap: make(map[uuid.UUID]User)}
}
func NewPGStorageUser(pool *pgxpool.Pool) *PostgresqlStorageUser {
	return &PostgresqlStorageUser{pool: pool}
}

func (m *MapStorageUser) GetAllUser() (map[uuid.UUID]User, error) {
	return m.storageUserMap, nil
}

func (p *PostgresqlStorageUser) GetAllUser() (map[uuid.UUID]User, error) {
	ctx := context.Background()
	sqlReq := `SELECT id, name, email FROM users`
	rows, err := p.pool.Query(ctx, sqlReq)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := make(map[uuid.UUID]User)

	for rows.Next() {
		var u User
		if err := rows.Scan(
			&u.ID,
			&u.Name,
			&u.Email,
		); err != nil {
			return nil, err
		}
		users[u.ID] = u
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (m *MapStorageUser) ExistEmailUser(email string) (bool, error) {
	for _, user := range m.storageUserMap {
		if user.Email == email {
			return true, nil
		}
	}
	return false, nil
}

func (p *PostgresqlStorageUser) ExistEmailUser(email string) (bool, error) {
	ctx := context.Background()
	var exist bool
	err := p.pool.QueryRow(ctx, `SELECT exists(SELECT 1 FROM users WHERE email=$1 )`, email).Scan(&exist)
	if err != nil {
		return false, err
	}
	return exist, nil
}

func (m *MapStorageUser) SaveUser(user User) error {
	m.storageUserMap[user.ID] = user
	return nil
}
func (p *PostgresqlStorageUser) SaveUser(user User) error {
	ctx := context.Background()
	sqlReq := `INSERT INTO users(
                  id,name,email,password) VALUES ($1,$2,$3,$4) ON CONFLICT (id) DO UPDATE SET
					name = EXCLUDED.name,
					email = EXCLUDED.email,
					password = EXCLUDED.password`
	_, err := p.pool.Exec(ctx, sqlReq, user.ID, user.Name, user.Email, user.Password)
	if err != nil {
		return err
	}
	return nil
}

func (m *MapStorageUser) GetUserID(id uuid.UUID) (*User, error) {
	user, _ := m.storageUserMap[id]
	return &user, nil
}

func (p *PostgresqlStorageUser) GetUserID(id uuid.UUID) (*User, error) {
	var user User
	ctx := context.Background()
	sqlReq := `SELECT id, name, email, password FROM users WHERE id=$1`
	err := p.pool.QueryRow(ctx, sqlReq, id).Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (m *MapStorageUser) ExistUser(id uuid.UUID) (bool, error) {
	_, exist := m.storageUserMap[id]
	if !exist {
		return false, nil
	}
	return true, nil
}

func (p *PostgresqlStorageUser) ExistUser(id uuid.UUID) (bool, error) {
	ctx := context.Background()
	var exist bool
	err := p.pool.QueryRow(ctx, `SELECT exists(SELECT 1 FROM users WHERE id=$1 )`, id).Scan(&exist)
	if err != nil {
		return false, err
	}
	return exist, nil
}

func (m *MapStorageUser) DeleteUser(id uuid.UUID) error {
	delete(m.storageUserMap, id)
	return nil
}

func (p *PostgresqlStorageUser) DeleteUser(id uuid.UUID) error {
	ctx := context.Background()
	_, err := p.pool.Exec(ctx, `DELETE FROM users WHERE id = $1`, id)
	if err != nil {
		return err
	}
	return nil
}
