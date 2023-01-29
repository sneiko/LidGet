package repository

import (
	"context"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5/pgxpool"
	"lidget/domain/entities"
	"lidget/domain/pkg/logger"
)

type RequestsRepository struct {
	db *pgxpool.Pool
}

func NewRequestsRepository(db *pgxpool.Pool) RequestsRepository {
	return RequestsRepository{db: db}
}

func (r *RequestsRepository) GetByRange(ctx context.Context, from int, to int) ([]*entities.Request, error) {
	if to == 0 {
		to = 500
	}

	count := to - from

	rows, err := r.db.Query(ctx, "select * from public.requests order by id desc limit $1 offset $2", count, from)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reqs []*entities.Request
	err = pgxscan.ScanAll(&reqs, rows)
	if err != nil {
		return nil, err
	}

	return reqs, nil
}

func (r *RequestsRepository) IsExistByText(ctx context.Context, text string) bool {
	sql := "select * from public.requests where text like $1"

	rows, err := r.db.Query(ctx, sql, text)
	if err != nil {
		logger.Error(err)
		return true
	}
	defer rows.Close()

	var reqs []*entities.Request
	err = pgxscan.ScanAll(&reqs, rows)
	if err != nil {
		logger.Error(err)
		return true
	}

	return len(reqs) > 0
}
