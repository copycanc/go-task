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
