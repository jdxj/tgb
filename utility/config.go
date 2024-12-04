package utility

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

func ParseCfgArray(ctx context.Context, pattern string, pointer any) error {
	v, err := g.Cfg().Get(ctx, pattern)
	if err != nil {
		return err
	}
	return gconv.Structs(v.Array(), pointer)
}

func ParseCfgInt64(ctx context.Context, pattern string) (int64, error) {
	v, err := g.Cfg().Get(ctx, pattern)
	if err != nil {
		return 0, err
	}
	return v.Int64(), nil
}
