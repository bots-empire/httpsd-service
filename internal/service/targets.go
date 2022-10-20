package service

import (
	"context"
	"github.com/pkg/errors"
	"httpsd-service/internal/entity"
)

func (m *Manager) AddTargetInDb(ctx context.Context, target *entity.TargetDb) error {
	m.logger.Info("targets")

	err := m.storage.AddTargetInDb(ctx, target)
	if err != nil {
		return errors.Wrap(err, "add target in db")
	}

	return nil
}

func (m *Manager) DeleteTargetFromDb(ctx context.Context, nameTarget string) error {
	m.logger.Info("targets")

	err := m.storage.DeleteTargetFromDb(ctx, nameTarget)
	if err != nil {
		errors.Wrap(err, "delete target in db")
	}

	return nil
}

func (m *Manager) GetTargetForPrometheus(ctx context.Context) ([]*entity.TargetPrometheus, error) {
	m.logger.Info("targets")
	targets, err := m.storage.GetTargetForPrometheus(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "get target in db")
	}

	return targets, nil
}
