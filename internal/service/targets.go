package service

import (
	"context"
	"encoding/json"
	"github.com/jackc/pgtype"
	"github.com/pkg/errors"
	"httpsd-service/internal/entity"
)

func (m *Manager) AddTargetInDb(ctx context.Context, target *entity.TargetDb) error {
	m.logger.Info("targets")

	targetDTO := limitToDTO(target)
	err := m.storage.AddTargetInDb(ctx, targetDTO)
	if err != nil {
		return errors.Wrap(err, "add target in db")
	}

	return nil
}

func (m *Manager) DeleteTargetFromDb(ctx context.Context, nameTarget string) error {
	return nil
}

func (m *Manager) GetTargetForPrometheus(ctx context.Context, nameTarget string) (*entity.TargetPrometheus, error) {
	return nil, nil
}

func limitToDTO(t *entity.TargetDb) *entity.TargetDbDTO {

	lbls, _ := json.Marshal(t.Labels)

	return &entity.TargetDbDTO{
		IpAddress: pgtype.Name{
			String: t.IpAddress,
			Status: pgtype.Present,
		},
		Labels: pgtype.JSONB{
			Bytes:  lbls,
			Status: pgtype.Present,
		},
	}
}
