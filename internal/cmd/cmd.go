package cmd

import (
	"context"

	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/jdxj/tgb/internal/controller"
)

var Main = gcmd.Command{
	Name:  "tgb",
	Usage: "tgb",
	Brief: "start tgb",
	Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
		return controller.Run(ctx)
	},
}
