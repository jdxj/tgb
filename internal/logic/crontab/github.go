package crontab

import (
	"context"
	"fmt"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtimer"
	"github.com/jdxj/tgb/internal/model"
	"github.com/jdxj/tgb/internal/service"
	"github.com/jdxj/tgb/utility"
)

type job struct {
	name    string
	pattern string
	f       gtimer.JobFunc
}

func hasNewTagJob() job {
	return job{
		name:    "HasNewTag",
		pattern: "0 0 9-19 * * 1-5",
		f: func(ctx context.Context) {
			var repos []*model.Repository
			err := utility.ParseCfgArray(ctx, "github.repos", &repos)
			if err != nil {
				g.Log().Errorf(ctx, "ParseCfgArray err: %s", err)
				return
			}

			for _, repo := range repos {
				tag, isNew, err := service.Github().LatestTag(ctx, repo)
				if err != nil {
					g.Log().Errorf(ctx, "LatestTag err: %s", err)
					continue
				}
				if !isNew {
					continue
				}

				content := fmt.Sprintf("repo [%s] new tag [%s]", repo.Id(), tag)
				err = service.Bot().Send(ctx, content)
				if err != nil {
					g.Log().Errorf(ctx, "Send err: %s", err)
				}
			}
		},
	}
}
