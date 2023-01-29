package repository

import (
	"context"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5/pgxpool"
	"lidget/domain/entities"
)

type CategoriesRepository struct {
	db *pgxpool.Pool
}

func NewCategoriesRepository(db *pgxpool.Pool) CategoriesRepository {
	return CategoriesRepository{db: db}
}

func (r *CategoriesRepository) GetAll(ctx context.Context) ([]*entities.Category, error) {
	rows, err := r.db.Query(ctx, "select * from public.category")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reqs []*entities.Category
	err = pgxscan.ScanAll(&reqs, rows)
	if err != nil {
		return nil, err
	}

	return reqs, nil
}
