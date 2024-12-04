package bot

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/jdxj/tgb/internal/service"
	tele "gopkg.in/telebot.v4"
)

func init() {
	service.RegisterBot(New())
}

type sBot struct {
	b *tele.Bot
}

func New() *sBot {
	ctx := gctx.GetInitCtx()
	pref := tele.Settings{
		Token: g.Cfg().MustGet(ctx, "tgb.token").String(),
		Poller: &tele.Webhook{
			Listen: ":8080",
			Endpoint: &tele.WebhookEndpoint{
				PublicURL: g.Cfg().MustGet(ctx, "tgb.url").String(),
			},
		},
	}

	bot, err := tele.NewBot(pref)
	if err != nil {
		g.Log().Fatalf(ctx, "NewBot err: %s", err)
	}
	return &sBot{b: bot}
}

func (s *sBot) Start(ctx context.Context) error {
	s.RegisterHandler(ctx)
	s.b.Start()
	return nil
}

func (s *sBot) Stop(ctx context.Context) error {
	s.b.Stop()
	return nil
}

func (s *sBot) RegisterHandler(ctx context.Context) error {
	s.b.Handle(hello(ctx))
	return nil
}
