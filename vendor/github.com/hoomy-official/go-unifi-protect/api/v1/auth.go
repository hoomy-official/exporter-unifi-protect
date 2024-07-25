package v1

import (
	"context"

	"github.com/hoomy-official/go-shared/pkg/net/do"
	"github.com/hoomy-official/go-shared/pkg/net/rest"
)

const (
	loginPath = "/api/auth/login"
)

type APIAuth struct {
	cl rest.Requester
}

func NewAPIAuth(cl rest.Requester) *APIAuth {
	return &APIAuth{cl: cl}
}

type LoginRequest struct {
	Password   interface{} `json:"password"`
	RememberMe bool        `json:"rememberMe"`
	Token      string      `json:"token"`
	Username   interface{} `json:"username"`
}

func (receiver *APIAuth) Login(ctx context.Context, request LoginRequest, authUser *User) error {
	return receiver.cl.POST(ctx, do.WithPath(loginPath), do.WithMarshalBody(request), do.WithUnmarshalBody(authUser))
}
