package utility

import (
	"testing"

	"github.com/gogf/gf/v2/os/gctx"
)

var ctx = gctx.GetInitCtx()

func TestParseCfgArray(t *testing.T) {
	type repo struct {
		Owner string
		Name  string
	}
	var res []repo
	err := ParseCfgArray(ctx, "github.repos", &res)
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	for _, v := range res {
		t.Logf("%+v\n", v)
	}
}

func TestParseCfgInt64(t *testing.T) {
	n, err := ParseCfgInt64(ctx, "tgb.owner")
	if err != nil {
		t.Logf("%s", err)
	}
	t.Log(n)
}
