package main

import (
	"github.com/gogf/gf/v2/os/gctx"

	"github.com/jdxj/tgb/internal/cmd"
	_ "github.com/jdxj/tgb/internal/logic"
	_ "github.com/jdxj/tgb/internal/packed"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
