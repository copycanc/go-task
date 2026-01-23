package db

import (
	"context"
	"fmt"
	"go-br-task/internal/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPostgresPool(cfg *config.PostgresCfg) (*pgxpool.Pool, error) {
	strConnect := fmt.Sprintf("postgres://%s:%s@%s:5432/%s?sslmode=disable", cfg.User, cfg.Pass, cfg.Host, cfg.DB)

	ctx := context.Background()
	//Подключаемся к БД (создаём pool)
	pool, err := pgxpool.New(ctx, strConnect)
	if err != nil {
		return nil, err
	}
	//Проверяем соединение
	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, err
	}
	return pool, nil
}

func CreateTable(pool *pgxpool.Pool) error {
	ctx := context.Background()
	createTask := `CREATE TABLE IF NOT EXISTS tasks (
					id   UUID PRIMARY KEY,
					title TEXT NOT NULL,
					description TEXT NOT NULL,
					status TEXT NOT NULL,
					created_at TIMESTAMP WITH TIME ZONE NOT NULL,
                    completed_at TIMESTAMP WITH TIME ZONE             
					)`
	createUser := `CREATE TABLE IF NOT EXISTS users (
					id   UUID PRIMARY KEY,
					name TEXT NOT NULL,
					email TEXT NOT NULL,
					password TEXT NOT NULL         
					)`
	_, err := pool.Exec(ctx, createTask)
	if err != nil {
		return err
	}
	_, err = pool.Exec(ctx, createUser)
	if err != nil {
		return err
	}
	return nil
}
