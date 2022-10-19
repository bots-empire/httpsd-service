package v1

import (
	"context"
	"encoding/json"
	"fmt"
	"httpsd-service/internal/entity"
	"httpsd-service/internal/service"
	"net/http"

	"go.uber.org/zap"
)

func HandleRouts(mux *http.ServeMux, m *service.Manager, logger *zap.Logger) {
	mux.HandleFunc("/v1/targets/add", func(w http.ResponseWriter, req *http.Request) {
		targ, err := targetsFromDbRequest(req)
		if err != nil {
			logger.Warn("error parse entity", zap.Any("err", err))
			http.Error(w, fmt.Sprintf("failed to parse entity: %v", err), http.StatusUnprocessableEntity)
			return
		}

		err = m.AddTargetInDb(context.Background(), targ)
		if err != nil {
			logger.Warn("error add target", zap.Any("err", err))
			http.Error(w, fmt.Sprintf("failed add target: %v", err), http.StatusInternalServerError)
			return
		}
	})
	logger.Sugar().Info("handle rout: /v1/targets/add")
	//одна ручка для добавления таргетов в бд
	//одна ручка для удаления данных из бд
	//однв ручка для забора данных для прометей из бд
}

func targetsFromDbRequest(req *http.Request) (*entity.TargetDb, error) {
	decoder := json.NewDecoder(req.Body)
	var a entity.TargetDb
	err := decoder.Decode(&a)
	if err != nil {
		return nil, err
	}

	return &a, nil
}
