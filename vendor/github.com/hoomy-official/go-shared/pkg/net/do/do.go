package do

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"go.uber.org/zap"
)

// Do sends an HTTP request using the provided context and URL with
// the configured options. It returns the HTTP response and any error
// encountered. It first applies the defaultOptions followed by any
// user-provided options. It then constructs and executes the request
// while logging the process at various stages using pre-request and
// post-request handlers.
func Do(ctx context.Context, u *url.URL, options ...Option) error {
	defaultOptions := []Option{
		WithMethod(http.MethodGet),
		WithClient(http.DefaultClient),
		WithLogger(zap.NewNop()),
		WithNow(time.Now),
	}

	p := NewParams()

	// Apply the default options followed by any user-provided options.
	for _, option := range append(defaultOptions, options...) {
		option(p)
	}

	// Record the start time for logging purposes.
	start := p.now()
	log := p.Logger.With(zap.Time("start", start))

	// Build the HTTP request with the applied options.
	log.Debug("buildRequest", zap.Duration("duration", time.Since(start)))
	req, err := http.NewRequestWithContext(ctx, p.Method, u.JoinPath(p.Path).String(), p.Body)
	if err != nil {
		log.Error("cannot buildRequest", zap.Error(err))
		return err
	}

	// Execute any pre-request handlers before sending the request.
	for name, preRequestHandler := range p.PreRequestHandler {
		log.Debug("preRequest", zap.Duration("duration", time.Since(start)), zap.String("preRequestHandlerName", name))
		if err = preRequestHandler.Apply(ctx, req); err != nil {
			log.Error("cannot handle request", zap.Error(err), zap.String("preRequestHandlerName", name))
			return err
		}
	}

	// Send the request.
	log.Debug("sendRequest", zap.Duration("duration", time.Since(start)))
	res, err := p.Client.Do(req)
	if err != nil {
		log.Error("cannot sendRequest", zap.Error(err))
		return err
	}

	defer res.Body.Close() // OK

	// Execute any post-request handlers after receiving the response.
	for name, postRequestHandler := range p.PostRequestHandler {
		log.Debug("postRequest", zap.Duration("duration", time.Since(start)), zap.String("postRequestHandlerName", name))

		if err = postRequestHandler.Apply(ctx, req, res); err != nil {
			log.Error("cannot handle response", zap.Error(err), zap.String("postRequestHandlerName", name))
			return err
		}
	}

	return nil
}
