package bot

import (
	"context"
	"fmt"

	tele "gopkg.in/telebot.v4"

	"github.com/jdxj/tgb/internal/service"
)

type handler struct {
	e any
	f tele.HandlerFunc
}

func (s *sBot) registerHandler(ctx context.Context) {
	handlers := []handler{
		helloHandler(),
		trafficHandler(ctx),
	}
	for _, handler := range handlers {
		s.b.Handle(handler.e, handler.f)
	}
}

func helloHandler() handler {
	return handler{
		e: "/hello",
		f: func(c tele.Context) error {
			sender := c.Sender()
			content := fmt.Sprintf("Hello %s%s!", sender.FirstName, sender.LastName)
			return c.Send(content)
		},
	}
}

func trafficHandler(ctx context.Context) handler {
	return handler{
		e: "/traffic",
		f: func(c tele.Context) error {
			err := service.Kiwi().YesterdayTraffic(ctx)
			return err
		},
	}
}
