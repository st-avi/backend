package utility

import (
	"errors"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func migrateConnect() *migrate.Migrate {
	m, err := migrate.New(
		"file://./manifest/migrations/postgresql/",
		"postgres://"+
			DBDefaultCfg.User+":"+DBDefaultCfg.Pass+"@"+
			DBDefaultCfg.Host+":"+DBDefaultCfg.Port+"/"+
			DBDefaultCfg.Name,
	)
	if err != nil {
		panic(err)
	}
	return m
}

func MigrateUp() {
	m := migrateConnect()
	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		panic(err)
	}
}

func MigrateDown() {
	m := migrateConnect()
	if err := m.Down(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		panic(err)
	}
}
