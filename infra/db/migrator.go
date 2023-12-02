package db

import (
	"context"
	"fmt"

	"github.com/Xilesun/sheethub/migrations"
	"github.com/uptrace/bun/migrate"
)

// IMigrator is the interface of the database migrator.
type IMigrator interface {
	Up(ctx context.Context) error
	Down(ctx context.Context) error
}

// Migrator is the type of the database migrator.
type Migrator struct {
	ctx context.Context
	*migrate.Migrator
}

// NewMigrator creates a new database migrator.
func NewMigrator(db *DB) (*Migrator, error) {
	migrator := migrate.NewMigrator(db.Client, migrations.Migrations)
	err := migrator.Init(db.ctx)
	if err != nil {
		return nil, err
	}
	return &Migrator{
		ctx:      db.ctx,
		Migrator: migrator,
	}, nil
}

// Up migrates the database up.
func (m *Migrator) Up() error {
	if err := m.Lock(m.ctx); err != nil {
		return err
	}
	defer m.Unlock(m.ctx)

	group, err := m.Migrate(m.ctx)
	if err != nil {
		return err
	}
	if group.IsZero() {
		fmt.Printf("there are no new migrations to run (database is up to date)\n")
		return nil
	}
	fmt.Printf("migrated to %s\n", group)
	return nil
}

// Down migrates the database down.
func (m *Migrator) Down() error {
	if err := m.Lock(m.ctx); err != nil {
		return err
	}
	defer m.Unlock(m.ctx)

	group, err := m.Rollback(m.ctx)
	if err != nil {
		return err
	}
	if group.IsZero() {
		fmt.Printf("there are no groups to roll back\n")
		return nil
	}
	fmt.Printf("rolled back %s\n", group)
	return nil
}
