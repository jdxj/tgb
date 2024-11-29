package cmd

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
	tele "gopkg.in/telebot.v4"
)

var Main = gcmd.Command{
	Name:  "main",
	Usage: "main",
	Brief: "start http server",
	Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
		pref := tele.Settings{
			Token: g.Cfg().MustGet(ctx, "tgb.token").String(),
			Poller: &tele.Webhook{
				Listen:            ":8080",
				MaxConnections:    0,
				AllowedUpdates:    []string{},
				IP:                "",
				DropUpdates:       false,
				SecretToken:       "",
				IgnoreSetWebhook:  false,
				HasCustomCert:     false,
				PendingUpdates:    0,
				ErrorUnixtime:     0,
				ErrorMessage:      "",
				SyncErrorUnixtime: 0,
				Endpoint: &tele.WebhookEndpoint{
					PublicURL: g.Cfg().MustGet(ctx, "tgb.url").String(),
				},
			},
		}

		b, err := tele.NewBot(pref)
		if err != nil {
			return err
		}

		b.Handle("/hello", func(c tele.Context) error {
			return c.Send("Hello!")
		})

		b.Start()
		return nil
	},
}
