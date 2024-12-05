package crontab

import (
	"context"
	"fmt"

	"github.com/gogf/gf/v2/frame/g"

	"github.com/jdxj/tgb/internal/service"
)

func yesterdayTrafficJob() job {
	return job{
		name:    "YesterdayTraffic",
		pattern: "0 5 9 * * *",
		f: func(ctx context.Context) {
			traffic, err := service.Kiwi().YesterdayTraffic(ctx)
			if err != nil {
				g.Log().Errorf(ctx, "YesterdayTraffic err: %s", err)
				return
			}

			content := fmt.Sprintf("kiwi yesterday traffic: %.3fGB", float64(traffic)/1000/1000/1000)
			err = service.Bot().Send(ctx, content)
			if err != nil {
				g.Log().Errorf(ctx, "Send err: %s", err)
			}
		},
	}
}
