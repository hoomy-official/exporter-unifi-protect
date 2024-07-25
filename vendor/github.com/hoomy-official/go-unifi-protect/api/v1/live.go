package v1

import (
	"context"
	"net/http"
	"net/url"

	"github.com/hoomy-official/go-shared/pkg/net/ws"
	"golang.org/x/sync/errgroup"
)

type APILive struct {
	ws ws.Dialer
}

func NewAPILive(cl ws.Dialer) *APILive {
	return &APILive{ws: cl}
}

type UpdateRequest struct {
	LastUpdateID string
}

func (receiver *APILive) Updates(ctx context.Context, r UpdateRequest, ch chan *Message) (*http.Response, error) {
	u := url.URL{Path: "/proxy/protect/ws/updates"}

	if r.LastUpdateID != "" {
		u.Query().Set("lastUpdateId", r.LastUpdateID)
	}

	conn, res, err := receiver.ws.Dial(ctx, u.String())
	if err != nil {
		return res, err
	}

	group, ctx := errgroup.WithContext(ctx)
	group.Go(func() error {
		for {
			_, i, er := conn.Read(ctx)
			if er != nil {
				return er
			}

			message, er := DecodeWsMessage(i)
			if er != nil {
				return er
			}

			ch <- message
		}
	})

	return nil, group.Wait()
}
