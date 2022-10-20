package targets

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"httpsd-service/internal/entity"
)

type Implementation interface {
	AddTargetInDb(ctx context.Context, target *entity.TargetDb) error
	DeleteTargetFromDb(ctx context.Context, nameTarget string) error

	GetTargetForPrometheus(ctx context.Context) ([]*entity.TargetPrometheus, error)
}

type Storage struct {
	db *pgxpool.Pool
}

func NewStorage(connect *pgxpool.Pool) *Storage {
	return &Storage{
		db: connect,
	}
}
