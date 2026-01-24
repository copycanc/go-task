package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPostgresPool() (*pgxpool.Pool, error) {
	uriStr := os.Getenv("DATABASE_URL")
	strConnect := fmt.Sprintf(uriStr)

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
