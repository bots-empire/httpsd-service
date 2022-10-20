package service

import (
	"go.uber.org/zap"
	"httpsd-service/internal/db/targets"
)

type Manager struct {
	storage targets.Implementation

	whiteList []int64

	logger *zap.Logger
}

func NewManager(logger *zap.Logger, storage targets.Implementation, whiteList []int64) *Manager {
	return &Manager{
		storage:   storage,
		whiteList: whiteList,
		logger:    logger,
	}
}
