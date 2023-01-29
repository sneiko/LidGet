package entities

import (
	"RxListener/pkg/logger"
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

type RequestCategory struct {
	Id         int `db:"id"`
	RequestId  int `db:"request_id"`
	CategoryId int `db:"category_id"`
}

func (r *RequestCategory) Create(ctx context.Context, db *pgxpool.Pool) (*int, error) {
	sql := `INSERT INTO public.request_category (id, request_id, category_id) 
			VALUES (DEFAULT, $1, $2) RETURNING id`

	var id int
	err := db.QueryRow(ctx, sql, r.RequestId, r.CategoryId).Scan(&id)
	if err != nil {
		return nil, err
	}

	logger.Info(sql)

	return &id, nil
}
