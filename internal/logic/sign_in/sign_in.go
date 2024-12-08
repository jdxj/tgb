package sign_in

import (
	"context"
	"net/http"

	"github.com/go-resty/resty/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/jdxj/tgb/internal/model"
	"github.com/jdxj/tgb/internal/service"
	"github.com/jdxj/tgb/utility"
)

func init() {
	service.RegisterSignIn(New())
}

func New() *sSignIn {
	return &sSignIn{
		hc: resty.New(),
	}
}

type sSignIn struct {
	hc *resty.Client
}

func (s *sSignIn) SignIn(ctx context.Context) error {
	var cookies []model.Cookie
	err := utility.ParseCfgArray(ctx, "sign_in", &cookies)
	if err != nil {
		return err
	}

	for _, cookie := range cookies {
		hCookies, err := http.ParseCookie(cookie.Cookie)
		if err != nil {
			return err
		}
		_, err = s.hc.R().SetCookies(hCookies).
			SetHeader("User-Agent", cookie.UA).
			Get(cookie.URL)
		if err != nil {
			g.Log().Errorf(ctx, "Get err: %s, name: %s", err, cookie.Name)
		}
	}
	return nil
}
