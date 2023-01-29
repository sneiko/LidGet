package entities

import (
	"RxListener/pkg/logger"
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

type Request struct {
	Id         int        `db:"id" json:"id"`
	SenderTgId string     `db:"sender_tg_id" json:"senderTgId"`
	Text       string     `db:"text" json:"text"`
	CreatedAt  time.Time  `db:"created_at" json:"createdAt"`
	UpdatedAt  *time.Time `db:"updated_at" json:"updatedAt"`
	DeletedAt  *time.Time `db:"deleted_at" json:"deletedAt"`
}

func (e *Request) Create(ctx context.Context, db *pgxpool.Pool) (*int, error) {
	sql := `INSERT INTO public.requests (id, sender_tg_id, text, created_at, updated_at, deleted_at) 
			VALUES (DEFAULT, $1, $2, now(), null, null) RETURNING id`

	var id int
	err := db.QueryRow(ctx, sql, e.SenderTgId, e.Text).Scan(&id)
	if err != nil {
		return nil, err
	}

	logger.Info(sql)

	return &id, nil
}
