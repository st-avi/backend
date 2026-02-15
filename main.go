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
	migrate := flag.String("migrate", "", "database migrate command: up / down")
	flag.Parse()

	switch *migrate {
	case "up":
		utility.MigrateUp()
		fmt.Println("Migration up completed.")
		return
	case "down":
		utility.MigrateDown()
		fmt.Println("Migration down completed.")
		return
	}

	cmd.Main.Run(gctx.GetInitCtx())
}
