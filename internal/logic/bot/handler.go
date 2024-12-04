package bot

import (
	"fmt"

	tele "gopkg.in/telebot.v4"
)

type handler struct {
	path string
	f    tele.HandlerFunc
}

func helloHandler() handler {
	return handler{
		path: "/hello",
		f: func(c tele.Context) error {
			sender := c.Sender()
			content := fmt.Sprintf("Hello %s!", sender.Username)
			return c.Send(content)
		},
	}
}
