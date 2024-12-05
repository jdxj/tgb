package crontab

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"

	"github.com/jdxj/tgb/internal/service"
)

func hasNewTagJob() job {
	return job{
		name:    "HasNewTag",
		pattern: "0 0 9-19 * * 1-5",
		f: func(ctx context.Context) {
			err := service.Github().HasNewTag(ctx)
			if err != nil {
				g.Log().Errorf(ctx, "HasNewTag err: %s", err)
			}
		},
	}
}

func yesterdayTrafficJob() job {
	return job{
		name:    "YesterdayTraffic",
		pattern: "0 5 9 * * *",
		f: func(ctx context.Context) {
			err := service.Kiwi().YesterdayTraffic(ctx)
			if err != nil {
				g.Log().Errorf(ctx, "YesterdayTraffic err: %s", err)
			}
		},
	}
}