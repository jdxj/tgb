package bot

import (
	"context"
	"fmt"

	tele "gopkg.in/telebot.v4"
)

func hello(_ context.Context) (path string, handler tele.HandlerFunc) {
	path = "/hello"
	handler = func(c tele.Context) error {
		sender := c.Sender()
		content := fmt.Sprintf("Hello %s!", sender.Username)
		return c.Send(content)
	}
	return
}
