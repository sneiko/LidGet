package repository

import (
	"context"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5/pgxpool"
	"lidget/domain/entities"
)

type CategoryPatternRepository struct {
	db *pgxpool.Pool
}

func NewCategoryPatternRepository(db *pgxpool.Pool) CategoryPatternRepository {
	return CategoryPatternRepository{db: db}
}

func (r *CategoryPatternRepository) GetAll(ctx context.Context) ([]entities.CategoryPattern, error) {
	sql := "select * from public.category_patterns"
	rows, err := r.db.Query(ctx, sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []entities.CategoryPattern
	err = pgxscan.ScanAll(&result, rows)
	if err != nil {
		return nil, err
	}

	return result, nil
}
