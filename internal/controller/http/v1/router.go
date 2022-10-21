package v1

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"

	"httpsd-service/internal/entity"
	"httpsd-service/internal/service"

	"go.uber.org/zap"
)

func HandleRouts(mux *http.ServeMux, m *service.Manager, logger *zap.Logger) {
	mux.HandleFunc("/v1/targets/add", func(w http.ResponseWriter, req *http.Request) {
		targ, err := targetFromRequest(req)
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

	mux.HandleFunc("/v1/targets/delete", func(w http.ResponseWriter, req *http.Request) {
		targ, err := targetFromRequest(req)
		if err != nil {
			logger.Warn("error parse entity", zap.Any("err", err))
			http.Error(w, fmt.Sprintf("failed to parse entity: %v", err), http.StatusUnprocessableEntity)
			return
		}

		err = m.DeleteTargetFromDb(context.Background(), targ.IpAddress)
		if err != nil {
			logger.Warn("error add target", zap.Any("err", err))
			http.Error(w, fmt.Sprintf("failed add target: %v", err), http.StatusInternalServerError)
			return
		}
	})
	logger.Sugar().Info("handle rout: /v1/targets/delete")

	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		targets, err := m.GetTargetForPrometheus(context.Background())
		if err != nil {
			logger.Warn("error add target", zap.Any("err", err))
			http.Error(w, fmt.Sprintf("failed add target: %v", err), http.StatusInternalServerError)
			return
		}

		data, err := targetsToJSON(targets)
		if err != nil {
			logger.Warn("error marshal to json", zap.Any("err", err))
			http.Error(w, fmt.Sprintf("failed marshal to json: %v", err), http.StatusUnprocessableEntity)
			return
		}

		w.Header().Set("Content-type", "application/json")
		_, err = w.Write(data)
		if err != nil {
			logger.Warn("error write array of bytes", zap.Any("err", err))
			http.Error(w, fmt.Sprintf("failed write array of bytes: %v", err), http.StatusUnprocessableEntity)
			return
		}
	})
	logger.Sugar().Info("handle rout: /")
}

func targetsToJSON(targets []*entity.TargetPrometheus) ([]byte, error) {
	data, err := json.Marshal(targets)
	if err != nil {
		return nil, errors.Wrap(err, "failed marshal targets")
	}

	return data, nil
}

func targetFromRequest(req *http.Request) (*entity.TargetDb, error) {
	decoder := json.NewDecoder(req.Body)
	var a entity.TargetDb
	err := decoder.Decode(&a)
	if err != nil {
		return nil, err
	}

	return &a, nil
}
