package main

import (
	_ "github.com/jdxj/tgb/internal/packed"

	_ "github.com/jdxj/tgb/internal/logic"

	"github.com/gogf/gf/v2/os/gctx"

	"github.com/jdxj/tgb/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
