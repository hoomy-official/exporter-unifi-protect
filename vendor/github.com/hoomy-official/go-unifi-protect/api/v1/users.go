package v1

import (
	"context"

	"github.com/hoomy-official/go-shared/pkg/net/do"
	"github.com/hoomy-official/go-shared/pkg/net/rest"
)

type APIUser struct {
	cl rest.Requester
}

func NewAPIUsers(cl rest.Requester) *APIUser {
	return &APIUser{cl: cl}
}

func (receiver *APIUser) Self(ctx context.Context, userUser *User) error {
	return receiver.cl.GET(ctx, do.WithPath("/api/users/self"), do.WithUnmarshalBody(userUser))
}
