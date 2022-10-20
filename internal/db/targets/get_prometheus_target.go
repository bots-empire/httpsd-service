package targets

import (
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
	"httpsd-service/internal/entity"
)

const GetTargetsForPrometheusFromDb = `SELECT ip_address, labels FROM http_sd.targets;`

func (s *Storage) GetTargetForPrometheus(ctx context.Context) ([]*entity.TargetPrometheus, error) {
	rows, err := s.db.Query(
		ctx,
		GetTargetsForPrometheusFromDb,
	)
	if err != nil {
		return nil, errors.Wrap(err, "bad request to get target for prometheus")
	}

	targets, err := readTargetsRows(rows)

	return targets, nil
}

func readTargetsRows(rows pgx.Rows) ([]*entity.TargetPrometheus, error) {
	var targets []*entity.TargetPrometheus

	for rows.Next() {
		target := &entity.TargetPrometheus{}

		if err := rows.Scan(
			&target.Targets,
			&target.Labels,
		); err != nil {
			return nil, errors.Wrap(err, "scan rows read targets")
		}

		targets = append(targets, target)
	}

	return targets, nil
}
