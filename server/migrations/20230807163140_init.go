package migrations

import (
	"context"

	"github.com/Xilesun/sheethub/server/models"
	"github.com/uptrace/bun"
)

func init() {
	Migrations.MustRegister(func(ctx context.Context, db *bun.DB) error {
		db.RegisterModel((*models.Sheet)(nil), (*models.Field)(nil))
		return nil
	}, func(ctx context.Context, db *bun.DB) error {
		return nil
	})
}
