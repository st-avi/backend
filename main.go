package main

import (
	"backend/internal/cmd"
	_ "backend/internal/packed"
	"backend/utility"
	"flag"
	"fmt"

	_ "github.com/gogf/gf/contrib/drivers/pgsql/v2"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	"github.com/gogf/gf/v2/os/gctx"
)

func main() {
	var up, down bool
	var version int
	flag.BoolVar(&up, "up", false, "migrate up to newest")
	flag.BoolVar(&down, "down", false, "migrate down to oldest")
	flag.IntVar(&version, "force", 0, "force set migration version")
	flag.Parse()

	if up || down || version > 0 {
		migration := utility.MigrateNew()
		if up {
			migration.MigrateUp()
			fmt.Println("Migration up completed.")
			return
		} else if down {
			migration.MigrateDown()
			fmt.Println("Migration down completed.")
			return
		} else if version > 0 {
			migration.Force(version)
			fmt.Println("Migration force completed.")
			return
		}
	}

	cmd.Main.Run(gctx.GetInitCtx())
}
