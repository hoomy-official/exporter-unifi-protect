package pkg

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/hoomy-official/go-shared/pkg/net/do"
	"github.com/hoomy-official/go-shared/pkg/net/rest"
	"github.com/hoomy-official/go-shared/pkg/net/ws"
	v1 "github.com/hoomy-official/go-unifi-protect/api/v1"
	"go.uber.org/zap"
)

type API interface {
	V1() *v1.V1
}
type APIClient struct {
	rest.Requester
	ws.Dialer

	httpClient *http.Client
	auth       *Auth

	v1 *v1.V1
}

func (c *APIClient) V1() *v1.V1 {
	return c.v1
}

func NewClient(baseURL *url.URL, auth *Auth, logger *zap.Logger, options ...Option) API {
	apiClient := &APIClient{}
	auth.baseURL = baseURL

	for _, option := range append([]Option{WithDefaultHTTPClient()}, options...) {
		option.apply(apiClient)
	}

	apiClient.Requester = rest.NewRest(
		baseURL,
		do.WithJSONRequest(),
		do.WithLogger(logger),
		do.WithClient(apiClient.httpClient),
		WithDefaultHTTPErrorCodeHandler(),
		WithRestAuthAdapter(auth),
	)

	apiClient.Dialer = NewDialer(baseURL, WithAuth(auth))
	apiClient.v1 = v1.NewV1(apiClient, apiClient)

	return apiClient
}

type ErrorMessage struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Level   string `json:"level"`
}

func WithRestAuthAdapter(auth *Auth) do.Option {
	return do.WithPreRequestHandler("a", func(ctx context.Context, r *http.Request) error {
		return auth.UserAuth(ctx, r)
	})
}

func WithDefaultHTTPErrorCodeHandler() do.Option {
	return do.WithPostRequestHandler("http_response_errorCode_handler", defaultRequestHandler)
}

func defaultRequestHandler(_ context.Context, _ *http.Request, response *http.Response) error {
	if response.StatusCode >= http.StatusBadRequest {
		return fmt.Errorf(
			"unexcpected response status code %d: %s",
			response.StatusCode,
			http.StatusText(response.StatusCode),
		)
	}
	return nil
}
