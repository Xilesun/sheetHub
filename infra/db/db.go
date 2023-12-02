package db

import (
	"context"
	"database/sql"

	"github.com/Xilesun/sheethub/infra/config"
	"github.com/Xilesun/sheethub/infra/constants"
	"github.com/Xilesun/sheethub/infra/errs"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/sqliteshim"
	"github.com/uptrace/bun/schema"
)

// DB is the implementation of the database.
type DB struct {
	ctx      context.Context
	Client   *bun.DB
	Migrator *Migrator
}

// SetupDB sets up the database.
func SetupDB(ctx context.Context, config config.DBConfig) (*DB, error) {
	db := &DB{
		ctx: ctx,
	}
	err := db.Connect(config)
	if err != nil {
		return nil, err
	}
	db.Migrator, err = NewMigrator(db)
	return db, err
}

func (db *DB) getDSN(config config.DBConfig) (dsn string, driver string, dialect schema.Dialect, err error) {
	switch config.Dialect {
	case constants.DialectSQLite:
		driver := sqliteshim.ShimName
		dialect := sqlitedialect.New()
		if config.DSN == "" {
			return "", driver, dialect, errs.New(errs.ErrDBConnect, "DSN is required for SQLite")
		}
		return config.DSN, driver, dialect, nil
	default:
		return "", "", nil, errs.New(errs.ErrDBConnect, "Unsupported dialect")
	}
}

// Connect connects to the database.
func (db *DB) Connect(config config.DBConfig) error {
	dsn, driver, dialect, err := db.getDSN(config)
	if err != nil {
		return err
	}
	sqldb, err := sql.Open(driver, dsn)
	if err != nil {
		return err
	}
	db.Client = bun.NewDB(sqldb, dialect)
	return nil
}
