package pkg

import (
	"context"
	"io"
	"net/http"
	"net/url"

	"github.com/hoomy-official/go-shared/pkg/net/ws"
	"nhooyr.io/websocket"
)

type Dialer struct {
	baseURL *url.URL
	params  *Params
}

func NewDialer(baseURL *url.URL, opts ...WSOption) *Dialer {
	params := &Params{}

	for _, opt := range opts {
		opt.Apply(params)
	}

	return &Dialer{baseURL: baseURL, params: params}
}

func (d Dialer) Dial(ctx context.Context, u string) (ws.Connecter, *http.Response, error) {
	h := d.params.HTTPHeader

	if h == nil {
		h = http.Header{}
	}

	r := http.Request{
		Header: h,
	}

	if d.params.auth != nil {
		// this is not nice, put dial options does not accept a Doer (only a plain HTTPClient)
		err := d.params.auth.UserAuth(ctx, &r)
		if err != nil {
			return nil, nil, err
		}
	}

	c, res, err := websocket.Dial(ctx, d.baseURL.JoinPath(u).String(), &websocket.DialOptions{
		HTTPClient:           d.params.HTTPClient,
		HTTPHeader:           r.Header,
		Host:                 d.params.Host,
		Subprotocols:         d.params.Subprotocols,
		CompressionMode:      websocket.CompressionMode(d.params.CompressionMode),
		CompressionThreshold: d.params.CompressionThreshold,
	})
	if err != nil {
		return nil, res, err
	}

	conn := &Connection{
		conn: c,
	}

	return conn, res, err
}

type Connection struct {
	conn *websocket.Conn
}

func (d *Connection) Reader(ctx context.Context) (ws.MessageType, io.Reader, error) {
	messageType, reader, err := d.conn.Reader(ctx)
	return ws.MessageType(messageType), reader, err
}

func (d *Connection) Read(ctx context.Context) (ws.MessageType, []byte, error) {
	typ, r, err := d.Reader(ctx)
	if err != nil {
		return 0, nil, err
	}

	b, err := io.ReadAll(r)
	return typ, b, err
}

func (d *Connection) Close(code ws.StatusCode, reason string) error {
	return d.conn.Close(websocket.StatusCode(code), reason)
}

func (d *Connection) CloseNow() error {
	return d.conn.CloseNow()
}

// CompressionMode represents the modes available to the permessage-deflate extension.
// See https://tools.ietf.org/html/rfc7692
//
// Works in all modern browsers except Safari which does not implement the permessage-deflate extension.
//
// Compression is only used if the peer supports the mode selected.
type CompressionMode int

// Params represents Dial's options.
type Params struct {
	auth *Auth

	// HTTPClient is used for the connection.
	// Its Transport must return writable bodies for WebSocket handshakes.
	// http.Transport does beginning with Go 1.12.
	HTTPClient *http.Client

	// HTTPHeader specifies the HTTP headers included in the handshake request.
	HTTPHeader http.Header

	// Host optionally overrides the Host HTTP header to send. If empty, the value
	// of URL.Host will be used.
	Host string

	// Subprotocols lists the WebSocket subprotocols to negotiate with the server.
	Subprotocols []string

	// CompressionMode controls the compression mode.
	// Defaults to CompressionDisabled.
	//
	// See docs on CompressionMode for details.
	CompressionMode CompressionMode

	// CompressionThreshold controls the minimum size of a message before compression is applied.
	//
	// Defaults to 512 bytes for CompressionNoContextTakeover and 128 bytes
	// for CompressionContextTakeover.
	CompressionThreshold int
}

type WSOption func(params *Params)

func (receiver WSOption) Apply(params *Params) {
	receiver(params)
}

func WithAuth(auth *Auth) WSOption {
	return func(params *Params) {
		params.auth = auth
	}
}
