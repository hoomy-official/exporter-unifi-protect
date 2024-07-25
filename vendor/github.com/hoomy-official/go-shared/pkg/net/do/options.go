package do

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"time"

	"go.uber.org/zap"
)

// Params is a structure that holds all the information
// needed to make an HTTP request and process its response.
// It includes HTTP client, request details, pre- and post-request handlers,
// a logger, and a timing function for testing.
type Params struct {
	Client HTTPClientDoer

	Method string
	Path   string
	Body   io.Reader

	PreRequestHandler  map[string]PreRequestHandlerFunc
	PostRequestHandler map[string]PostRequestHandlerFunc

	Logger *zap.Logger

	now func() time.Time
}

// HTTPClientDoer performs a HTTP request.
type HTTPClientDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// NewParams initiates an instance of Params with default values.
// It sets up the maps for pre-request and post-request handlers.
func NewParams() *Params {
	return &Params{
		PreRequestHandler:  map[string]PreRequestHandlerFunc{},
		PostRequestHandler: map[string]PostRequestHandlerFunc{},
	}
}

// PreRequestHandlerFunc represents a function that is invoked before making
// an HTTP request. It has the chance to modify the request based on the context.
type PreRequestHandlerFunc func(ctx context.Context, res *http.Request) error

// Apply executes the PreRequestHandlerFunc on the given request, using the provided context.
func (receiver PreRequestHandlerFunc) Apply(ctx context.Context, req *http.Request) error {
	return receiver(ctx, req)
}

// PostRequestHandlerFunc represents a function that is invoked after receiving
// the HTTP response. It can be used to handle the response or its body accordingly.
type PostRequestHandlerFunc func(ctx context.Context, req *http.Request, res *http.Response) error

// Apply executes the PostRequestHandlerFunc on the given response, using the provided context.
func (receiver PostRequestHandlerFunc) Apply(ctx context.Context, req *http.Request, res *http.Response) error {
	return receiver(ctx, req, res)
}

// Option is a function type that modifies parameters for an HTTP request.
type Option func(params *Params)

// Apply executes the Option on the given Params.
func (p Option) Apply(params *Params) {
	p(params)
}

// WithLogger returns an option function that sets the logger in the Params.
func WithLogger(logger *zap.Logger) Option {
	return func(params *Params) {
		params.Logger = logger
	}
}

// WithNow returns an option function that sets the function to return current time
// for timing or testing purposes in the Params.
func WithNow(fn func() time.Time) Option {
	return func(params *Params) {
		params.now = fn
	}
}

// WithClient returns an option function that sets the HTTP client in the Params.
func WithClient(cl HTTPClientDoer) Option {
	return func(params *Params) {
		params.Client = cl
	}
}

// WithMethod returns an option function that sets the HTTP method for the request in the Params.
func WithMethod(method string) Option {
	return func(params *Params) {
		params.Method = method
	}
}

// WithExtraHeader returns an option function that sets one HTTP header for the request in the Params.
func WithExtraHeader(key, value string) Option {
	return WithPreRequestHandler(
		fmt.Sprintf("http_request_set_header_%s", key),
		func(_ context.Context, req *http.Request) error {
			req.Header.Set(key, value)
			return nil
		},
	)
}

// WithExtraHeaderf returns an option function that formats a value according to its format and sets one HTTP
// header for the request in the Params.
func WithExtraHeaderf(key, format string, a ...any) Option {
	return WithExtraHeader(key, fmt.Sprintf(format, a...))
}

// WithHeader returns an option function that sets the HTTP header for the request in the Params.
func WithHeader(header http.Header) Option {
	return WithPreRequestHandler(
		"http_request_set_header",
		func(_ context.Context, req *http.Request) error {
			for key, strings := range header {
				for _, str := range strings {
					req.Header.Set(key, str)
				}
			}
			return nil
		},
	)
}

// WithContentLength returns an option function that sets the Content length for the request in the Params.
func WithContentLength(requestContent []byte) Option {
	return WithPreRequestHandler(
		"http_request_content_length",
		func(_ context.Context, req *http.Request) error {
			req.ContentLength = int64(len(requestContent))
			return nil
		},
	)
}

// WithPath returns an option function that sets the request path in the Params.
// The path can be formatted with the provided arguments.
func WithPath(path string, a ...any) Option {
	return func(params *Params) {
		params.Path = fmt.Sprintf(path, a...)
	}
}

// WithBody returns an option function that sets the request body in the Params.
func WithBody(b io.Reader) Option {
	return func(params *Params) {
		params.Body = b
	}
}

// WithMarshalBody returns an option function that sets a pre-request handler
// to marshal the provided value into JSON and use it as the request body.
func WithMarshalBody(v any) Option {
	return WithPreRequestHandler(
		"http_request_body_json_unmarshal",
		func(_ context.Context, req *http.Request) error {
			b, err := json.Marshal(v)
			if err != nil {
				return err
			}

			req.Body = io.NopCloser(bytes.NewReader(b))

			return nil
		},
	)
}

// WithPreRequestHandler returns an option that associates the provided
// PreRequestHandlerFunc with the given name for use by Params.
func WithPreRequestHandler(name string, f PreRequestHandlerFunc) Option {
	return func(params *Params) {
		params.PreRequestHandler[name] = f
	}
}

// WithJSONRequest returns an option function that sets a pre-request handler
// in the Params to automatically set "Content-Type" header to "application/json".
func WithJSONRequest() Option {
	return WithPreRequestHandler(
		"http_request_header_json",
		func(_ context.Context, req *http.Request) error {
			req.Header.Set("Content-Type", "application/json")
			return nil
		},
	)
}

// WithPostRequestHandler returns an option that associates the provided
// PostRequestHandlerFunc with the given name for use by Params.
func WithPostRequestHandler(name string, f PostRequestHandlerFunc) Option {
	return func(params *Params) {
		params.PostRequestHandler[name] = f
	}
}

// WithUnmarshalBody returns an option function that sets a post-request handler
// to unmarshal the response body into the provided value.
func WithUnmarshalBody(v any) Option {
	return WithPostRequestHandler(
		"http_response_body_json_unmarshal",
		func(_ context.Context, _ *http.Request, res *http.Response) error {
			if v == nil || (reflect.ValueOf(v).Kind() == reflect.Ptr && reflect.ValueOf(v).IsNil()) {
				return nil
			}

			body, err := io.ReadAll(res.Body)
			if err != nil {
				return err
			}

			err = json.Unmarshal(body, v)
			if err != nil {
				return err
			}

			res.Body = io.NopCloser(bytes.NewBuffer(body))
			return nil
		},
	)
}
