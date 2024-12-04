package bot

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	tele "gopkg.in/telebot.v4"

	"github.com/jdxj/tgb/internal/service"
	"github.com/jdxj/tgb/utility"
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
	s.registerHandler()
	s.b.Start()
	return nil
}

func (s *sBot) Stop(ctx context.Context) error {
	s.b.Stop()
	return nil
}

func (s *sBot) registerHandler() {
	handlers := []handler{
		helloHandler(),
	}
	for _, handler := range handlers {
		s.b.Handle(handler.path, handler.f)
	}
}

func (s *sBot) Send(ctx context.Context, content string) error {
	id, err := utility.ParseCfgInt64(ctx, "tgb.owner")
	if err != nil {
		return err
	}
	user := &tele.User{ID: id}
	_, err = s.b.Send(user, content)
	return err
}
