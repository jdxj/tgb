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
	jobs := []job{
		hasNewTagJob(),
	}
	for _, job := range jobs {
		_, err := gcron.AddSingleton(ctx, job.pattern, job.f, job.name)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *sCrontab) Stop(ctx context.Context) error {
	gcron.Stop()
	return nil
}
