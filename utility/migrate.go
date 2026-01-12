package utility

import (
	"errors"
	"strings"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func migrateConnect() *migrate.Migrate {
	ctx := gctx.New()
	dbDefaultLink, _ := g.Cfg().Get(ctx, "database.default.link")
	m, err := migrate.New(
		"file://./manifest/migrations/postgresql/",
		strings.Replace(dbDefaultLink.String(), "pgsql:", "postgres://", 1),
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
