package service

import (
	"go.uber.org/zap"
	"httpsd-service/internal/db/metrics"
)

type Manager struct {
	storage metrics.Implementation

	whiteList []int64

	logger *zap.Logger
}

func NewManager(logger *zap.Logger, storage metrics.Implementation, whiteList []int64) *Manager {
	return &Manager{
		storage:   storage,
		whiteList: whiteList,
		logger:    logger,
	}
}
