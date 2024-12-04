package logic

import (
	"testing"

	"github.com/gogf/gf/v2/os/gctx"
	"github.com/jdxj/tgb/internal/model"
)

var (
	ctx = gctx.GetInitCtx()
	s   = New()
)

func TestLatestTag(t *testing.T) {
	repo := &model.Repository{
		Owner: "golang-design",
		Name:  "clipboard",
	}
	tag, _, err := s.LatestTag(ctx, repo)
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	t.Logf("tag: %s\n", tag)
}
