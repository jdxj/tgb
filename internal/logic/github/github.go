package logic

import (
	"context"
	"fmt"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/google/go-github/v67/github"

	"github.com/jdxj/tgb/internal/model"
	"github.com/jdxj/tgb/internal/service"
	"github.com/jdxj/tgb/utility"
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

func (s *sGithub) LatestTag(ctx context.Context, repo *model.Repository) (string, bool, error) {
	tags, _, err := s.c.Repositories.ListTags(ctx, repo.Owner, repo.Name, &github.ListOptions{Page: 1, PerPage: 1})
	if err != nil {
		return "", false, err
	}
	if len(tags) <= 0 {
		return "", false, nil
	}

	var (
		latestTag = tags[0].GetName()
		preTag    = s.tags[repo.Id()]
	)
	s.tags[repo.Id()] = latestTag
	if preTag == "" || preTag == latestTag {
		return latestTag, false, nil
	}
	return latestTag, true, nil
}

func (s *sGithub) HasNewTag(ctx context.Context) error {
	var repos []*model.Repository
	err := utility.ParseCfgArray(ctx, "github.repos", &repos)
	if err != nil {
		return err
	}

	for _, repo := range repos {
		tag, isNew, err := s.LatestTag(ctx, repo)
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
	return nil
}
