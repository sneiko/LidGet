package repository

import "github.com/jackc/pgx/v5/pgxpool"

type Repositories struct {
	Requests         RequestsRepository
	Categories       CategoriesRepository
	CategoryPatterns CategoryPatternRepository
}

func NewRepositories(db *pgxpool.Pool) Repositories {
	return Repositories{
		Requests:         NewRequestsRepository(db),
		Categories:       NewCategoriesRepository(db),
		CategoryPatterns: NewCategoryPatternRepository(db),
	}
}
