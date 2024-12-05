package crontab

import (
	"context"

	"github.com/gogf/gf/v2/os/gcron"
	"github.com/gogf/gf/v2/os/gtimer"

	"github.com/jdxj/tgb/internal/service"
)

func init() {
	service.RegisterCrontab(New())
}

type job struct {
	name    string
	pattern string
	f       gtimer.JobFunc
}

func New() *sCrontab {
	return new(sCrontab)
}

type sCrontab struct{}

func (s *sCrontab) Start(ctx context.Context) error {
	return s.registerJob(ctx)
}

func (s *sCrontab) Stop(ctx context.Context) error {
	gcron.Stop()
	return nil
}
