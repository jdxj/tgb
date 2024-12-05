// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
)

type (
	IKiwi interface {
		YesterdayTraffic(ctx context.Context) error
	}
)

var (
	localKiwi IKiwi
)

func Kiwi() IKiwi {
	if localKiwi == nil {
		panic("implement not found for interface IKiwi, forgot register?")
	}
	return localKiwi
}

func RegisterKiwi(i IKiwi) {
	localKiwi = i
}
