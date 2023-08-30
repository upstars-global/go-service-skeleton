package pgsql

import (
	"database/sql"
	"errors"
	"os"
	"path/filepath"

	"github.com/upstars-global/go-service-skeleton/pkg/config"
	"github.com/upstars-global/go-service-skeleton/pkg/logger"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type DB struct {
	*sql.DB

	cfg config.DBConfigProvider
	log logger.Interface
}

type DBTXer interface {
	Begin() (*sql.Tx, error)
}

func New(cfg config.DBConfigProvider, log logger.Interface) (db *DB, err error) {
	db = &DB{
		cfg: cfg,
		log: log,
	}
	if db.DB, err = sql.Open("postgres", cfg.GetDBDSN()); err != nil { // todo: apm
		log.With("err", err, "dsn", cfg.GetDBDSN()).Error("db: sql.Open() failed")
		return
	}
	db.DB.SetMaxOpenConns(cfg.GetDBMaxOpenCons())
	db.DB.SetMaxIdleConns(cfg.GetDBMaxIdleCons())
	db.log.Debug("db: connection established")
	err = db.migrate()
	return
}

func (db *DB) Rollback() (err error) {
	var driver database.Driver
	err = db.Ping()
	if err != nil {
		db.log.With("err", err).Error("db: Rollback(): db.Ping() failed")
		return
	}

	if driver, err = postgres.WithInstance(db.DB, &postgres.Config{}); err != nil {
		db.log.With("err", err).Error("db: Rollback(): postgres.WithInstance() failed")
		return
	}

	var p string
	p, _ = os.Getwd()
	p, _ = filepath.Abs(p + "/" + db.cfg.GetDBMigrations())

	var m *migrate.Migrate
	if m, err = migrate.NewWithDatabaseInstance("file:"+p, "", driver); err != nil {
		db.log.With("err", err).Error("db: Rollback(): migrate.NewWithDatabaseInstance() failed")
		return
	}

	if err = m.Down(); err != nil && !errors.Is(err, migrate.ErrNoChange) && !os.IsNotExist(err) {
		db.log.With("err", err).Error("db: Rollback(): m.Down() failed")
		return
	}
	err = nil

	db.log.Debug("db: Rollback(): Rollback proceed")
	return
}

func (db *DB) migrate() (err error) {
	var driver database.Driver

	err = db.Ping()
	if err != nil {
		db.log.With("err", err).Error("db: migrate(): db.Ping() failed")
		return
	}

	if driver, err = postgres.WithInstance(db.DB, &postgres.Config{}); err != nil {
		db.log.With("err", err).Error("db: migrate(): postgres.WithInstance() failed")
		return
	}

	var p string
	p, _ = os.Getwd()
	p, _ = filepath.Abs(p + "/" + db.cfg.GetDBMigrations())

	var m *migrate.Migrate
	if m, err = migrate.NewWithDatabaseInstance("file:"+p, "", driver); err != nil {
		db.log.With("err", err).Error("db: migrate(): migrate.NewWithDatabaseInstance() failed")
		return
	}

	if err = m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) && !os.IsNotExist(err) {
		db.log.With("err", err).Error("db: migrate(): m.Up() failed")
		return
	}
	err = nil

	db.log.Debug("db: migrate(): migrations proceed")
	return
}
