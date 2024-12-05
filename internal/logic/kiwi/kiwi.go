package logic

import (
	"context"
	"fmt"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gtime"
	kiwi "github.com/jdxj/kiwivm-api-go"

	"github.com/jdxj/tgb/internal/service"
)

func init() {
	service.RegisterKiwi(New())
}

func New() *sKiwi {
	ctx := gctx.GetInitCtx()
	veId := g.Cfg().MustGet(ctx, "kiwivm.veId").String()
	key := g.Cfg().MustGet(ctx, "kiwivm.key").String()
	s := &sKiwi{
		c: kiwi.NewClient(veId, key),
	}
	return s
}

type sKiwi struct {
	c *kiwi.Client
}

func (s *sKiwi) YesterdayTraffic(ctx context.Context) error {
	stats, err := s.c.GetRawUsageStats(ctx)
	if err != nil {
		return err
	}

	var (
		yesterday = gtime.Now().AddDate(0, 0, -1)
		beginUnix = yesterday.StartOfDay().Unix()
		endUnix   = yesterday.EndOfDay().Unix()
	)
	yesterdayTraffic := stats.Traffic(beginUnix, endUnix)
	content := fmt.Sprintf("kiwi yesterday traffic: %.3fGB", float64(yesterdayTraffic)/1000/1000/1000)
	return service.Bot().Send(ctx, content)
}
