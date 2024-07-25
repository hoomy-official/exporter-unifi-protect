package v1

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"

	"github.com/hoomy-official/go-shared/pkg/net/do"
	"github.com/hoomy-official/go-shared/pkg/net/rest"
)

const apiPath = "/proxy/protect/api/sensors"

type APISensor struct {
	cl rest.Requester
}

func NewAPISensors(cl rest.Requester) *APISensor {
	return &APISensor{
		cl: cl.New(do.WithPath(apiPath)),
	}
}

func (receiver *APISensor) List(ctx context.Context, sensors *[]Sensor) error {
	return receiver.cl.GET(ctx, WithUnmarshalBody(sensors))
}

func (receiver *APISensor) Get(ctx context.Context, id string, sensor *Sensor) error {
	return receiver.cl.GET(ctx, do.WithPath("%s", id), do.WithUnmarshalBody(sensor))
}

func WithUnmarshalBody(v any) do.Option {
	//nolint:bodyclose // not necessary, the body will reset
	return do.WithPostRequestHandler("http_response_body_json_unmarshal", unmarshalHandler(v))
}

func unmarshalHandler(v any) func(_ context.Context, _ *http.Request, response *http.Response) error {
	return func(_ context.Context, _ *http.Request, response *http.Response) error {
		if v == nil || (reflect.ValueOf(v).Kind() == reflect.Ptr && reflect.ValueOf(v).IsNil()) {
			return nil
		}

		body, err := io.ReadAll(response.Body)
		if err != nil {
			return err
		}

		if !json.Valid(body) {
			return fmt.Errorf("not a valid json: %w", err)
		}

		err = json.Unmarshal(body, v)
		if err != nil {
			return err
		}

		response.Body = io.NopCloser(bytes.NewBuffer(body))
		return nil
	}
}
