package do

import (
	"context"
	"net/url"
)

// Doer defines an interface for sending HTTP requests and handling their responses.
// The Do method sends an HTTP request with the provided context and any optional
// request options for adjusting the request prior to execution.
type Doer interface {
	// Do sends an HTTP request and return the error.
	Do(ctx context.Context, options ...Option) error
}

type D func(ctx context.Context, options ...Option) error

func (receiver D) Do(ctx context.Context, options ...Option) error {
	return receiver(ctx, options...)
}

func NewDoer(u *url.URL, baseOptions ...Option) Doer {
	return D(func(ctx context.Context, options ...Option) error {
		return Do(ctx, u, append(baseOptions, options...)...)
	})
}
