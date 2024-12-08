// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
)

type (
	ISignIn interface {
		SignIn(ctx context.Context) error
	}
)

var (
	localSignIn ISignIn
)

func SignIn() ISignIn {
	if localSignIn == nil {
		panic("implement not found for interface ISignIn, forgot register?")
	}
	return localSignIn
}

func RegisterSignIn(i ISignIn) {
	localSignIn = i
}
