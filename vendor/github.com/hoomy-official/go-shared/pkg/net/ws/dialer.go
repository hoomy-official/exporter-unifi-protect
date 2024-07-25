package ws

import (
	"context"
	"net/http"
)

// Dialer defines how to establish WebSocket connections.
type Dialer interface {
	// Dial creates and returns a new WebSocket connection.
	Dial(ctx context.Context, url string) (Connecter, *http.Response, error)
}
