package metrics

import (
	"context"
	"github.com/pkg/errors"
	"httpsd-service/internal/entity"
)

const addTargetQuery = `INSERT INTO http_sd.targerts (ip_address, labels)
VALUES ($1, convert_from($2,'utf8')::jsonb)
ON CONFLICT(ip_address) DO UPDATE SET labels = convert_from($2,'utf8')::jsonb`

func (s *Storage) AddTargetInDb(ctx context.Context, target *entity.TargetDbDTO) error {
	_, err := s.db.Exec(
		ctx,
		addTargetQuery,
		target.IpAddress,
		target.Labels,
	)
	if err != nil {
		return errors.Wrap(err, "bad request to add target in db")
	}

	return nil
}
