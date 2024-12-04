// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
)

type (
	IGithub interface {
		LatestTag(ctx context.Context, owner string, repo string) (string, bool, error)
	}
)

var (
	localGithub IGithub
)

func Github() IGithub {
	if localGithub == nil {
		panic("implement not found for interface IGithub, forgot register?")
	}
	return localGithub
}

func RegisterGithub(i IGithub) {
	localGithub = i
}
