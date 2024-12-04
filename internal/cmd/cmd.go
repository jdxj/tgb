package cmd

import (
	"context"

	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/jdxj/tgb/internal/service"
)

var Main = gcmd.Command{
	Name:  "tgb",
	Usage: "tgb",
	Brief: "start tgb",
	Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
		return service.Bot().Start(ctx)
	},
}
