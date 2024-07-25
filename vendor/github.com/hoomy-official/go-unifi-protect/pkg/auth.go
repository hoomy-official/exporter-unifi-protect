package pkg

import (
	"context"

	"github.com/golang-jwt/jwt/v5"
	"github.com/hoomy-official/go-shared/pkg/net/do"
	v1 "github.com/hoomy-official/go-unifi-protect/api/v1"

	"net/http"
	"net/url"
	"time"
)

const (
	TokenCookieName = "TOKEN"
	CSRFTokenHeader = "X-CSRF-Token" //nolint:gosec // not an issue
)

type Auth struct {
	baseURL *url.URL

	User *v1.User

	Username string
	Password string

	CSRF string

	Token  *jwt.Token
	Claims *jwt.RegisteredClaims
}

func NewAuth(username string, password string) *Auth {
	return &Auth{Username: username, Password: password, User: &v1.User{}, Claims: &jwt.RegisteredClaims{}}
}

func (a *Auth) UserAuth(ctx context.Context, r *http.Request) error {
	if a.Claims.ExpiresAt == nil || a.Claims.ExpiresAt.Before(time.Now()) {
		err := do.Do(
			ctx,
			a.baseURL,
			do.WithJSONRequest(),
			do.WithMethod(http.MethodPost),
			do.WithPath("/api/auth/login"),
			do.WithMarshalBody(v1.LoginRequest{Username: a.Username, Password: a.Password}),
			do.WithUnmarshalBody(a.User),
			do.WithPostRequestHandler("cookie_handler", a.userAuthResponseHandler),
		)
		if err != nil {
			return err
		}
	}

	r.Header.Set(CSRFTokenHeader, a.CSRF)
	r.AddCookie(&http.Cookie{Name: "TOKEN", Value: a.Token.Raw})

	return nil
}

func (a *Auth) userAuthResponseHandler(_ context.Context, _ *http.Request, res *http.Response) error {
	var err error

	for _, cookie := range res.Cookies() {
		if cookie.Name == TokenCookieName {
			a.Token, _, err = jwt.NewParser().ParseUnverified(cookie.Value, a.Claims)
			if err != nil {
				return err
			}

			continue
		}
	}

	a.CSRF = res.Header.Get(CSRFTokenHeader)

	return nil
}
