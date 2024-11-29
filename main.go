package main

import (
	_ "tgb/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"tgb/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
