package logic

import (
	"testing"

	"github.com/gogf/gf/v2/os/gctx"
)

var (
	ctx = gctx.GetInitCtx()
	s   = New()
)

func TestLatestTag(t *testing.T) {
	tag, _, err := s.LatestTag(ctx, "golang-design", "clipboard")
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	t.Logf("tag: %s\n", tag)
}
