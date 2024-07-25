package rest

import (
	"context"

	"github.com/hoomy-official/go-shared/pkg/net/do"
)

// Requester is an interface that defines the methods for making various types
// of HTTP requests such as GET, POST, PUT, and others.
type Requester interface {
	do.Doer

	// New create a new instance with new default options
	New(options ...do.Option) Requester

	GET(ctx context.Context, options ...do.Option) error
	POST(ctx context.Context, options ...do.Option) error
	PUT(ctx context.Context, options ...do.Option) error
	PATCH(ctx context.Context, options ...do.Option) error
	DELETE(ctx context.Context, options ...do.Option) error
	OPTIONS(ctx context.Context, options ...do.Option) error
	HEAD(ctx context.Context, options ...do.Option) error
	CONNECT(ctx context.Context, options ...do.Option) error
	TRACE(ctx context.Context, options ...do.Option) error
}
