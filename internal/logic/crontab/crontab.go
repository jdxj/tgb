package crontab

import (
	"context"

	"github.com/gogf/gf/v2/os/gcron"
)

func New() *sCrontab {
	return new(sCrontab)
}

type sCrontab struct{}

func (s *sCrontab) Start(ctx context.Context) error {
	gcron.AddSingleton()
}

func (s *sCrontab) Stop(ctx context.Context) error {
}
