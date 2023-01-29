package db

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"lidget/domain/config"
	"lidget/domain/pkg/logger"
)

func Connect(cfg *config.Config) (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(context.Background(), cfg.DbConnectionString)
	if err != nil {
		logger.Error("Failed to connect to database")
	}

	return pool, err
}
