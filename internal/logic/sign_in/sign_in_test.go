package sign_in

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/gogf/gf/v2/os/gctx"
	"github.com/jdxj/tgb/internal/model"
	"github.com/jdxj/tgb/utility"
)

var ctx = gctx.GetInitCtx()

func TestParseCookie(t *testing.T) {
	var cookies []model.Cookie
	err := utility.ParseCfgArray(ctx, "sign_in", &cookies)
	if err != nil {
		t.Fatal(err)
	}
	for _, cookie := range cookies {
		hCookies, err := http.ParseCookie(cookie.Cookie)
		if err != nil {
			t.Fatal(err)
		}
		for _, hCookie := range hCookies {
			fmt.Printf("%+v\n", hCookie)
		}
	}
}
