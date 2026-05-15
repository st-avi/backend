package utility

import (
	"errors"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type Migration struct {
	client *migrate.Migrate
}

func MigrateNew() *Migration {
	m := Migration{}
	var err error
	m.client, err = migrate.New(
		"file://./manifest/migrations/postgresql/",
		"postgres://"+
			DBDefaultCfg.User+":"+DBDefaultCfg.Pass+"@"+
			DBDefaultCfg.Host+":"+DBDefaultCfg.Port+"/"+
			DBDefaultCfg.Name,
	)
	if err != nil {
		panic(err)
	}
	return &m
}

func (m *Migration) MigrateUp() {
	if err := m.client.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		panic(err)
	}
}

func (m *Migration) MigrateDown() {
	if err := m.client.Down(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		panic(err)
	}
}

func (m *Migration) Force(version int) {
	if err := m.client.Force(version); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		panic(err)
	}
}
