package rest

import (
	"context"
	"net/http" // Correct the import path from "net/HTTP" to "net/http"
	"net/url"

	"github.com/hoomy-official/go-shared/pkg/net/do"
)

// Rest is a concrete implementation of the Requester interface.
// It uses a base URL and a set of options that can be customized for each request.
type Rest struct {
	baseOptions []do.Option
	baseURL     *url.URL
}

// NewRest creates a new Rest client with a given base URL and options.
func NewRest(baseURL *url.URL, options ...do.Option) *Rest {
	return &Rest{baseOptions: options, baseURL: baseURL}
}

// Do makes an HTTP request using the base URL and options of the Rest client
// along with additional options provided for the specific request.
func (r *Rest) Do(ctx context.Context, options ...do.Option) error {
	return do.Do(ctx, r.baseURL, append(r.baseOptions, options...)...)
}

// New create a new instance with new default options.
func (r *Rest) New(options ...do.Option) Requester {
	return &Rest{
		baseOptions: append(r.baseOptions, options...),
		baseURL:     r.baseURL,
	}
}

// GET performs an HTTP GET request with the specified options.
func (r *Rest) GET(ctx context.Context, options ...do.Option) error {
	return r.Do(ctx, append(options, do.WithMethod(http.MethodGet))...)
}

// POST performs an HTTP POST request with the specified options.
func (r *Rest) POST(ctx context.Context, options ...do.Option) error {
	return r.Do(ctx, append(options, do.WithMethod(http.MethodPost))...)
}

// PUT performs an HTTP PUT request with the specified options.
func (r *Rest) PUT(ctx context.Context, options ...do.Option) error {
	return r.Do(ctx, append(options, do.WithMethod(http.MethodPut))...)
}

// PATCH performs an HTTP PATCH request with the specified options.
func (r *Rest) PATCH(ctx context.Context, options ...do.Option) error {
	return r.Do(ctx, append(options, do.WithMethod(http.MethodPatch))...)
}

// DELETE performs an HTTP DELETE request with the specified options.
func (r *Rest) DELETE(ctx context.Context, options ...do.Option) error {
	return r.Do(ctx, append(options, do.WithMethod(http.MethodDelete))...)
}

// OPTIONS performs an HTTP OPTIONS request with the specified options.
func (r *Rest) OPTIONS(ctx context.Context, options ...do.Option) error {
	return r.Do(ctx, append(options, do.WithMethod(http.MethodOptions))...)
}

// HEAD performs an HTTP HEAD request with the specified options.
func (r *Rest) HEAD(ctx context.Context, options ...do.Option) error {
	return r.Do(ctx, append(options, do.WithMethod(http.MethodHead))...)
}

// CONNECT performs an HTTP CONNECT request with the specified options.
func (r *Rest) CONNECT(ctx context.Context, options ...do.Option) error {
	return r.Do(ctx, append(options, do.WithMethod(http.MethodConnect))...)
}

// TRACE performs an HTTP TRACE request with the specified options.
func (r *Rest) TRACE(ctx context.Context, options ...do.Option) error {
	return r.Do(ctx, append(options, do.WithMethod(http.MethodTrace))...)
}
