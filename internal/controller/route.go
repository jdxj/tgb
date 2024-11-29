package controller

import (
	"context"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	tele "gopkg.in/telebot.v4"
)

func Run(ctx context.Context) error {
	pref := tele.Settings{
		Token: g.Cfg().MustGet(ctx, "tgb.token").String(),
		Poller: &tele.Webhook{
			Listen: ":8080",
			Endpoint: &tele.WebhookEndpoint{
				PublicURL: g.Cfg().MustGet(ctx, "tgb.url").String(),
			},
		},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		return err
	}

	registerRoute(b)
	b.Start()
	return nil
}

func registerRoute(b *tele.Bot) {
	b.Handle("/date", func(c tele.Context) error {
		return c.Send(gtime.Now().Layout(time.DateTime))
	})
}
