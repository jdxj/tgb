// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
)

type (
	ICrontab interface {
		Start(ctx context.Context) error
		Stop(ctx context.Context) error
	}
)

var (
	localCrontab ICrontab
)

func Crontab() ICrontab {
	if localCrontab == nil {
		panic("implement not found for interface ICrontab, forgot register?")
	}
	return localCrontab
}

func RegisterCrontab(i ICrontab) {
	localCrontab = i
}
