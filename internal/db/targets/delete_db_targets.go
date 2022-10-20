package targets

import (
	"context"
	"github.com/pkg/errors"
)

const DeleteTargetFromDbQuery = `DELETE FROM http_sd.targets WHERE ip_address = $1;`

func (s *Storage) DeleteTargetFromDb(ctx context.Context, nameTarget string) error {
	_, err := s.db.Exec(
		ctx,
		DeleteTargetFromDbQuery,
		nameTarget,
	)
	if err != nil {
		return errors.Wrap(err, "bad request to delete from db")
	}

	return nil
}
