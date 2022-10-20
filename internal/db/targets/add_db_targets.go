package targets

import (
	"context"
	"encoding/json"
	"github.com/jackc/pgtype"
	"github.com/pkg/errors"
	"httpsd-service/internal/entity"
)

const addTargetQuery = `INSERT INTO http_sd.targerts (ip_address, labels)
VALUES ($1, convert_from($2,'utf8')::jsonb)
ON CONFLICT(ip_address) DO UPDATE SET labels = convert_from($2,'utf8')::jsonb;`

func (s *Storage) AddTargetInDb(ctx context.Context, target *entity.TargetDb) error {
	targetDTO := targetToDTO(target)

	_, err := s.db.Exec(
		ctx,
		addTargetQuery,
		targetDTO.IpAddress,
		targetDTO.Labels,
	)
	if err != nil {
		return errors.Wrap(err, "bad request to add target in db")
	}

	return nil
}

type targetDbDTO struct {
	IpAddress pgtype.Name  `json:"ip_address,omitempty"`
	Labels    pgtype.JSONB `json:"labels,omitempty"`
}

func targetToDTO(t *entity.TargetDb) *targetDbDTO {

	lbls, _ := json.Marshal(t.Labels)

	return &targetDbDTO{
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
