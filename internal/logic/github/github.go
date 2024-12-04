package logic

import (
	"context"
	"fmt"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/google/go-github/v67/github"
	"github.com/jdxj/tgb/internal/service"
)

func init() {
	service.RegisterGithub(New())
}

func New() *sGithub {
	s := &sGithub{
		c:    github.NewClient(nil).WithAuthToken(g.Cfg().MustGet(gctx.GetInitCtx(), "github.ghp").String()),
		tags: make(map[string]string),
	}
	return s
}

type sGithub struct {
	c *github.Client

	tags map[string]string
}

func repoId(owner, repo string) string {
	return fmt.Sprintf("%s/%s", owner, repo)
}

func (s *sGithub) LatestTag(ctx context.Context, owner, repo string) (string, bool, error) {
	tags, _, err := s.c.Repositories.ListTags(ctx, owner, repo, &github.ListOptions{Page: 1, PerPage: 1})
	if err != nil {
		return "", false, err
	}
	if len(tags) <= 0 {
		return "", false, nil
	}

	var (
		latestTag = tags[0].GetName()
		rid       = repoId(owner, repo)
		preTag    = s.tags[rid]
	)
	s.tags[rid] = latestTag
	if preTag == "" || preTag == latestTag {
		return latestTag, false, nil
	}
	return latestTag, true, nil
}
