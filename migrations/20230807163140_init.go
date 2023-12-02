package migrations

import (
	"context"
	"database/sql"

	"github.com/Xilesun/sheethub/models"
	"github.com/uptrace/bun"
)

func init() {
	Migrations.MustRegister(func(ctx context.Context, db *bun.DB) error {
		return db.RunInTx(ctx, &sql.TxOptions{}, func(ctx context.Context, tx bun.Tx) error {
			if _, err := db.NewCreateTable().Model((*models.Sheet)(nil)).Exec(ctx); err != nil {
				return err
			}
			_, err := db.NewCreateTable().Model((*models.Field)(nil)).Exec(ctx)
			return err
		})
	}, func(ctx context.Context, db *bun.DB) error {
		return nil
	})
}
