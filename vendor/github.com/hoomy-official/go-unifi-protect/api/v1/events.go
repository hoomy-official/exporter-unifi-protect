package v1

import (
	"context"

	"github.com/hoomy-official/go-shared/pkg/net/do"
	"github.com/hoomy-official/go-shared/pkg/net/rest"
)

type APIEvent struct {
	cl rest.Requester
}

func NewAPIEvents(cl rest.Requester) *APIEvent {
	return &APIEvent{cl: cl}
}

func (receiver *APIEvent) Self(ctx context.Context, _ LoginRequest, eventEvent *Event) error {
	return receiver.cl.GET(ctx, do.WithPath("/api/events/self"), do.WithUnmarshalBody(eventEvent))
}
